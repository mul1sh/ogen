package chain

import (
	"errors"
	"fmt"
	"sync"

	"github.com/olympus-protocol/ogen/p2p"
	"github.com/olympus-protocol/ogen/primitives"
	"github.com/olympus-protocol/ogen/txs/txverifier"
)

type txSchemes struct {
	Type   primitives.TxType
	Action primitives.TxAction
}

type TxPayloadInv struct {
	txs  map[txSchemes][]primitives.Tx
	lock sync.RWMutex
}

func (txp *TxPayloadInv) Add(scheme txSchemes, tx primitives.Tx, wg *sync.WaitGroup) {
	defer wg.Done()
	txp.lock.Lock()
	txp.txs[scheme] = append(txp.txs[scheme], tx)
	txp.lock.Unlock()
	return
}

var (
	ErrorTooManyGenerateTx = errors.New("chainProcessor-too-many-generate: the block contains more generate tx than expected")
	ErrorInvalidBlockSig   = errors.New("chainProcessor-block-sig-verify: the block signature verification failed")
	ErrorPubKeyNoMatch     = errors.New("chainProcessor-invalid-signer: the block signer is not valid")
)

func (ch *Blockchain) newTxPayloadInv(txs []primitives.Tx, blocks int) (*TxPayloadInv, error) {
	txPayloads := &TxPayloadInv{
		txs: make(map[txSchemes][]primitives.Tx),
	}
	var wg sync.WaitGroup
	for _, tx := range txs {
		wg.Add(1)
		scheme := txSchemes{
			Type:   tx.TxType,
			Action: tx.TxAction,
		}
		go func(scheme txSchemes, tx primitives.Tx) {
			txPayloads.Add(scheme, tx, &wg)
		}(scheme, tx)
	}
	wg.Wait()
	if len(txPayloads.txs[txSchemes{
		Type:   primitives.TxCoins,
		Action: primitives.Generate,
	}]) > blocks {
		return nil, ErrorTooManyGenerateTx
	}
	return txPayloads, nil
}

func (ch *Blockchain) ProcessBlockInv(blockInv p2p.MsgBlockInv) error {
	// TODO: this is disabled for now because we don't have transaction execution done.
	// if we have a block that spends an input, we need to update our state representation
	// for that block before we try to verify other blocks.

	//txs := blockInv.GetTxs()
	//txPayloadInv, err := ch.newTxPayloadInv(txs, len(blockInv.GetBlocks()))
	//if err != nil {
	//	return err
	//}
	//err = ch.verifyTx(txPayloadInv)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (ch *Blockchain) ProcessBlock(block *primitives.Block) error {
	// 1. first verify basic block properties
	// b. get parent block

	// 2. verify block against previous block's state
	oldState, found := ch.state.GetStateForHash(block.Header.PrevBlockHash)
	if !found {
		return fmt.Errorf("missing parent block state: %s", block.Header.PrevBlockHash)
	}

	txPayloadInv, err := ch.newTxPayloadInv(block.Txs, 1)
	if err != nil {
		ch.log.Warn(err)
		return err
	}

	// a. verify transactions
	ch.log.Debugf("tx inventory created types to verify: %v", len(txPayloadInv.txs))
	err = ch.verifyTx(oldState, txPayloadInv)
	if err != nil {
		ch.log.Warn(err)
		return err
	}
	ch.log.Debugf("tx verification finished successfully")

	// b. apply block transition to state
	ch.log.Debugf("attempting to apply block to state")
	// TODO: better fork choice here
	_, err = ch.State().Add(block, true)
	if err != nil {
		ch.log.Warn(err)
		return err
	}
	ch.log.Infof("New block accepted Hash: %v", block.Hash())

	return nil
}

func (ch *Blockchain) verifyTx(prevState *primitives.State, inv *TxPayloadInv) error {

	for scheme, txs := range inv.txs {
		txVerifier := txverifier.NewTxVerifier(&*prevState, &ch.params)
		err := txVerifier.VerifyTxsBatch(txs, scheme.Type, scheme.Action)
		if err != nil {
			return err
		}
	}
	return nil
}
