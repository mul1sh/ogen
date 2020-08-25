package bls

import (
	"github.com/olympus-protocol/ogen/pkg/bls/bls_herumi"
	"github.com/olympus-protocol/ogen/pkg/bls/bls_kilic"
	bls_interface "github.com/olympus-protocol/ogen/pkg/bls/interface"
	"github.com/olympus-protocol/ogen/pkg/params"
)

var CurrImplementation Implementation = bls_herumi.HerumiImplementation{}

// Initialize overrides current implementation with a new implementation
func Initialize(p params.ChainParams, lib string) error {
	switch lib {
	case "herumi":
		CurrImplementation = bls_herumi.HerumiImplementation{}
	case "kilic":
		CurrImplementation = bls_kilic.KilicImplementation{}
	default:
		CurrImplementation = bls_herumi.HerumiImplementation{}
	}
	bls_interface.Prefix = p.AccountPrefixes
	return nil
}

type Implementation interface {
	SignatureFromBytes(sig []byte) (bls_interface.Signature, error)
	NewAggregateSignature() bls_interface.Signature
	AggregateSignatures(sigs []bls_interface.Signature) bls_interface.Signature
	Aggregate(sigs []bls_interface.Signature) bls_interface.Signature
	RandKey() bls_interface.SecretKey
	SecretKeyFromBytes(privKey []byte) (bls_interface.SecretKey, error)
	PublicKeyFromBytes(pubKey []byte) (bls_interface.PublicKey, error)
}
