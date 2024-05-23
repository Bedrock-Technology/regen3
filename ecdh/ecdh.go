package ecdh

import (
	"crypto/ecdsa"
	"math/big"
)

func GetSecret(publicKey *ecdsa.PublicKey, key *ecdsa.PrivateKey) *big.Int {
	secret, _ := key.Curve.ScalarMult(publicKey.X, publicKey.Y, key.D.Bytes())
	return secret
}
