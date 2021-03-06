package burnproof

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/olympus-protocol/ogen/pkg/chainhash"
)

var merkleRootHash [32]byte

func init() {
	hashBytes, _ := hex.DecodeString("0be71bd3e3ec9046901b21f066407f8413c73dc3145d6a515d0fa03b28e0140f")
	copy(merkleRootHash[:], hashBytes)
}

// CoinsProof is a proof of coins on the old blockchain.
type CoinsProof struct {
	MerkleIndex  uint32
	MerkleBranch []chainhash.Hash
	PkScript     []byte
	Transaction  wire.MsgTx
}

func (c *CoinsProof) Marshal(w io.Writer) error {
	return nil
}

// Unmarshal decodes the proof from a byte slice.
func (c *CoinsProof) Unmarshal(r io.Reader) error {
	indexBytes := make([]byte, 4)
	if _, err := io.ReadFull(r, indexBytes); err != nil {
		return err
	}

	c.MerkleIndex = binary.LittleEndian.Uint32(indexBytes)

	merkleCount, err := wire.ReadVarInt(r, 0)
	if err != nil {
		return err
	}

	c.MerkleBranch = make([]chainhash.Hash, merkleCount)
	for i := range c.MerkleBranch {
		_, err := io.ReadFull(r, c.MerkleBranch[i][:])
		if err != nil {
			return err
		}
	}

	if err := c.Transaction.BtcDecode(r, 0, wire.BaseEncoding); err != nil {
		return err
	}

	script, err := wire.ReadVarBytes(r, 0, 10000, "script")
	if err != nil {
		return err
	}
	c.PkScript = script

	return nil
}

var cache = txscript.NewSigCache(1000)

// CalcMerkleRoot calculates the merkle root under some assumptions.
func (c *CoinsProof) CalcMerkleRoot(leafHash chainhash.Hash) chainhash.Hash {
	index := c.MerkleIndex
	for _, h := range c.MerkleBranch {
		if index&1 > 0 {
			leafHash = chainhash.DoubleHashH(append(h[:], leafHash[:]...))
		} else {
			leafHash = chainhash.DoubleHashH(append(leafHash[:], h[:]...))
		}
		index >>= 1
	}

	return leafHash
}

func calcLeafHash(proof *CoinsProof) (*chainhash.Hash, error) {
	if len(proof.Transaction.TxOut) != 1 {
		return nil, fmt.Errorf("expected transaction to have 1 output, but got %d", len(proof.Transaction.TxOut))
	}

	if len(proof.Transaction.TxIn) != 1 {
		return nil, fmt.Errorf("expected transaction to have 1 output, but got %d", len(proof.Transaction.TxOut))
	}

	leafHashBytes := bytes.NewBuffer([]byte{})
	outpoint := proof.Transaction.TxIn[0].PreviousOutPoint

	_, err := leafHashBytes.Write(outpoint.Hash[:])
	if err != nil {
		return nil, err
	}

	indexBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(indexBytes, outpoint.Index)

	_, err = leafHashBytes.Write(indexBytes)
	if err != nil {
		return nil, err
	}

	valueBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valueBytes, uint64(proof.Transaction.TxOut[0].Value))

	_, err = leafHashBytes.Write(valueBytes)
	if err != nil {
		return nil, err
	}

	if err := wire.WriteVarBytes(leafHashBytes, 0, proof.PkScript); err != nil {
		return nil, err
	}

	h := chainhash.DoubleHashH(leafHashBytes.Bytes())
	return &h, nil
}

func verifyMerkleRoot(merkleRoot chainhash.Hash, proof *CoinsProof) error {
	out, err := calcLeafHash(proof)
	if err != nil {
		return err
	}

	calculatedRoot := proof.CalcMerkleRoot(*out)

	if !calculatedRoot.IsEqual(&merkleRoot) {
		return fmt.Errorf("calculated merkle root %s, but expected %s", calculatedRoot, merkleRoot)
	}

	return nil
}

func verifyScript(proof *CoinsProof) error {
	if len(proof.Transaction.TxOut) != 1 {
		return fmt.Errorf("expected transaction to have 1 output, but got %d", len(proof.Transaction.TxOut))
	}

	eng, err := txscript.NewEngine(proof.PkScript, &proof.Transaction, 0,
		txscript.StandardVerifyFlags, cache,
		txscript.NewTxSigHashes(&proof.Transaction), proof.Transaction.TxOut[0].Value)
	if err != nil {
		return err
	}

	if err := eng.Execute(); err != nil {
		return err
	}

	return nil
}

func verifyPkhMatchesAddress(script []byte, address string) error {
	if len(script) != 25 {
		return fmt.Errorf("expected transaction pkscript to be 25, but got %d", len(script))
	}

	addrHash := script[3:23]

	addrBuf := new(bytes.Buffer)
	addrBytes := []byte(address)

	if err := wire.WriteVarInt(addrBuf, 0, uint64(len(addrBytes))); err != nil {
		return err
	}

	if _, err := addrBuf.Write(addrBytes); err != nil {
		return err
	}

	expectedAddrHash := chainhash.DoubleHashB(addrBuf.Bytes())
	if !bytes.Equal(addrHash, expectedAddrHash[:20]) {
		return fmt.Errorf("expected addresses to match (expected: %x, got: %x)", expectedAddrHash, addrHash)
	}

	return nil
}

// VerifyBurn verifies a burn proof.
func VerifyBurn(proofBytes []byte, address string) error {
	var proofs []*CoinsProof

	buf := bytes.NewBuffer(proofBytes)

	for {
		c := new(CoinsProof)
		err := c.Unmarshal(buf)
		if err != nil {
			return err
		}

		proofs = append(proofs, c)
		if buf.Len() <= 0 {
			break
		}
	}

	for _, c := range proofs {
		err := VerifyBurnProof(c, address)
		if err != nil {
			return err
		}
	}

	return nil
}

// VerifyBurnProof verifies a single burn proof.
func VerifyBurnProof(p *CoinsProof, address string) error {
	if err := verifyMerkleRoot(merkleRootHash, p); err != nil {
		return err
	}

	if err := verifyScript(p); err != nil {
		return err
	}

	if err := verifyPkhMatchesAddress(p.Transaction.TxOut[0].PkScript, address); err != nil {
		return err
	}

	return nil
}
