package indexer

import (
	"context"
	"encoding/hex"
	"github.com/olympus-protocol/ogen/api/proto"
	"github.com/olympus-protocol/ogen/cmd/ogen/indexer/db"
	"github.com/olympus-protocol/ogen/pkg/logger"
	"github.com/olympus-protocol/ogen/pkg/params"
	"github.com/olympus-protocol/ogen/pkg/primitives"
	"github.com/olympus-protocol/ogen/pkg/rpcclient"
	"io"
	"os"
	"sync"
)

// Indexer is the module that allows operations across multiple services.
type Indexer struct {
	log logger.Logger
	ctx context.Context

	client    *rpcclient.Client
	db        *db.Database
	canClose  *sync.WaitGroup
	netParams *params.ChainParams
}

func (i *Indexer) Start() {
	i.initialSync()
	i.log.Info("Listening for new blocks")
	go i.subscribeBlocks()
	i.mempoolSync()
	go i.subscribeMempool()
}

func (i *Indexer) subscribeBlocks() {
	subscribe, err := i.client.Chain().SubscribeBlocks(context.Background(), &proto.Empty{})
	if err != nil {
		panic("unable to initialize subscription client")
	}
	for {
		select {
		case <-i.ctx.Done():
			subscribe.CloseSend()
			break
		default:
			res, err := subscribe.Recv()
			if err == io.EOF || err != nil {
				// listener closed restart with sync
				i.initialSync()
				continue
			}
			// To make sure the explorer is always synced, every new block we reinsert the last 5
			blockBytes, err := hex.DecodeString(res.Data)
			if err != nil {
				i.log.Errorf("unable to parse error %s", err.Error())
				continue
			}
			block := new(primitives.Block)
			err = block.Unmarshal(blockBytes)
			if err != nil {
				i.log.Errorf("unable to parse error %s", err.Error())
				continue
			}
			err = i.db.InsertBlock(block)
			if err != nil {
				if err == db.ErrorPrevBlockHash {
					i.log.Error(db.ErrorPrevBlockHash)
					i.log.Info("Restarting sync...")
					i.initialSync()
					continue
				}
				i.log.Errorf("unable to insert error %s", err.Error())
				continue
			}
			i.log.Infof("Received new block %s", block.Hash().String())
		}
	}
}

func (i *Indexer) subscribeMempool() {
	syncClient, err := i.client.Utils().SubscribeMempool(i.ctx, &proto.Empty{})
	if err != nil {
		i.log.Fatal(err)
	}
	for {
		select {
		case <-i.ctx.Done():
			syncClient.CloseSend()
			break
		default:
			tx, err := syncClient.Recv()
			if err != nil {
				continue
			}
			i.log.Info("Received mempool transaction, indexing...")
			i.db.ProcessMempoolTransaction(tx)
		}

	}
}

func (i *Indexer) mempoolSync() {
	syncClient, err := i.client.Utils().SyncMempool(i.ctx, &proto.Empty{})
	if err != nil {
		i.log.Fatal(err)
	}
	for {
		tx, err := syncClient.Recv()
		if err != nil {
			if err == io.EOF {
				_ = syncClient.CloseSend()
				break
			}
			i.log.Error(err)
			break
		}
		i.db.ProcessMempoolTransaction(tx)
	}
}

func (i *Indexer) initialSync() {

	// get the saved state
	indexState, err := i.db.GetCurrentState()
	if err != nil {
		i.log.Fatal(err)
	}

	var latestBHash string
	if indexState.Blocks == 0 && indexState.LastBlockHash == "" {
		// Initialize the database with initialization parameters
		latestBHash, err = i.db.Initialize()
		if err != nil {
			i.log.Fatal("unable to initialize database")
			return
		}
	} else {
		latestBHash = indexState.LastBlockHash
	}

	i.log.Infof("Starting initial sync...")
	syncClient, err := i.client.Chain().Sync(context.Background(), &proto.Hash{Hash: latestBHash})
	if err != nil {
		i.log.Fatal("unable to initialize initial sync")
		return
	}

	blockCount := 0
	for {
		res, err := syncClient.Recv()
		if err != nil {
			if err == io.EOF {
				_ = syncClient.CloseSend()
				break
			}
			i.log.Error(err)
			break
		}
		blockBytes, err := hex.DecodeString(res.Data)
		if err != nil {
			i.log.Error("unable to parse block")
			break
		}
		block := new(primitives.Block)
		err = block.Unmarshal(blockBytes)
		if err != nil {
			i.log.Error("unable to parse block")
			break
		}
		err = i.db.InsertBlock(block)
		if err != nil {
			i.log.Error("unable to insert block")
			break
		} else {
			blockCount++
		}
	}
	i.log.Infof("Initial sync finished, parsed %d blocks", blockCount)
}

func (i *Indexer) Close() {
	i.db.Close()
}

func (i *Indexer) Context() context.Context {
	return i.ctx
}

func NewIndexer(dbConnString, rpcEndpoint, dbDriver string, netParams *params.ChainParams) (*Indexer, error) {
	log := logger.New(os.Stdin)

	rpcClient := rpcclient.NewRPCClient(rpcEndpoint, true)
	var wg sync.WaitGroup
	database := db.NewDB(dbConnString, log, &wg, dbDriver, netParams)

	err := database.Migrate()
	if err != nil {
		return nil, err
	}

	indexer := &Indexer{
		log:       log,
		ctx:       context.Background(),
		client:    rpcClient,
		db:        database,
		canClose:  &wg,
		netParams: netParams,
	}

	return indexer, nil
}
