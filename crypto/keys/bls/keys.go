// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package bls

import (
	"bytes"

	"github.com/minio/sha256-simd"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/crypto/tmhash"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/itsdevbear/comet-bls12-381/bls/blst"
	"github.com/itsdevbear/comet-bls12-381/bls/params"
)

const (
	// PrivKeySize defines the length of the PrivKey byte array.
	PrivKeySize = 32
	// PubKeySize defines the length of the PubKey byte array.
	PubKeySize = 48
	// KeyType is the string constant for the bls12_381 algorithm.
	KeyType = "bls12_381"
)

// -------------------------------------.
const (
	PrivKeyName = "tendermint/PrivKeyBLS12_381"
	PubKeyName  = "tendermint/PubKeyBLS12_381"
)

func init() {
	cmtjson.RegisterType(PubKey{}, PubKeyName)
	cmtjson.RegisterType(PrivKey{}, PrivKeyName)
}

// ===============================================================================================
// Private Key
// ===============================================================================================

// PrivKey is a wrapper around the Ethereum bls12_381 private key type. This wrapper conforms to
// crypotypes.Pubkey to allow for the use of the Ethereum bls12_381 private key type within the
// Cosmos SDK.

// Compile-time type assertion.
var _ cryptotypes.LedgerPrivKey = &PrivKey{}

func NewPrivateKeyFromBytes(bz []byte) (*PrivKey, error) {
	secretKey, err := blst.SecretKeyFromBytes(bz)
	if err != nil {
		return nil, err
	}
	return &PrivKey{secretKey.Marshal()}, nil
}

func GenPrivKey() (PrivKey, error) {
	secretKey, err := blst.RandKey()
	return PrivKey{secretKey.Marshal()}, err
}

// Bytes returns the byte representation of the ECDSA Private Key.
func (privKey PrivKey) Bytes() []byte {
	return privKey.Key
}

// PubKey returns the ECDSA private key's public key. If the privkey is not valid
// it returns a nil value.
func (privKey PrivKey) PubKey() cryptotypes.PubKey {
	secretKey, _ := blst.SecretKeyFromBytes(privKey.Key)
	key := PubKey{secretKey.PublicKey().Marshal()}
	return &key
}

// Equals returns true if two ECDSA private keys are equal and false otherwise.
func (privKey PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	return privKey.Type() == other.Type() && bytes.Equal(privKey.Bytes(), other.Bytes())
}

// Type returns eth_bls12_381.
func (privKey PrivKey) Type() string {
	return KeyType
}

func (privKey PrivKey) Sign(digestBz []byte) ([]byte, error) {
	secretKey, err := blst.SecretKeyFromBytes(privKey.Key)
	if err != nil {
		return nil, err
	}

	bz := digestBz
	if len(bz) > 32 {
		hash := sha256.Sum256(bz)
		bz = hash[:]
	}

	sig := secretKey.Sign(bz)
	return sig.Marshal(), nil
}

// ===============================================================================================
// Public Key
// ===============================================================================================

// Pubkey is a wrapper around the Ethereum bls12_381 public key type. This wrapper conforms to
// crypotypes.Pubkey to allow for the use of the Ethereum bls12_381 public key type within the
// Cosmos SDK.

// Compile-time type assertion.
var _ cryptotypes.PubKey = &PubKey{}

// Address returns the address of the ECDSA public key.
// The function will return an empty address if the public key is invalid.
func (pubKey PubKey) Address() crypto.Address {
	pk, _ := blst.PublicKeyFromBytes(pubKey.Key)
	if len(pk.Marshal()) != PubKeySize {
		panic("pubkey is incorrect size")
	}
	// TODO: do we want to keep this address format?
	return crypto.Address(tmhash.SumTruncated(pubKey.Key))
}

func (pubKey PubKey) VerifySignature(msg, sig []byte) bool {
	if len(sig) != params.BLSSignatureLength {
		return false
	}
	bz := msg
	if len(msg) > 32 {
		hash := sha256.Sum256(msg)
		bz = hash[:]
	}

	pubK, _ := blst.PublicKeyFromBytes(pubKey.Key)
	ok, err := blst.VerifySignature(sig, [32]byte(bz[:32]), pubK)
	if err != nil {
		return false
	}
	return ok
}

// Bytes returns the pubkey byte format.
func (pubKey PubKey) Bytes() []byte {
	return pubKey.Key
}

func (pubKey PubKey) String() string {
	return string(pubKey.Bytes())
}

// Type returns eth_bls12_381.
func (pubKey PubKey) Type() string {
	return KeyType
}

// Equals returns true if the pubkey type is the same and their bytes are deeply equal.
func (pubKey PubKey) Equals(other cryptotypes.PubKey) bool {
	return pubKey.Type() == other.Type() && bytes.Equal(pubKey.Bytes(), other.Bytes())
}
