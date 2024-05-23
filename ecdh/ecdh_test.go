package ecdh

import (
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSecret(t *testing.T) {
	key1, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	assert.Nil(t, err)
	key2, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	assert.Nil(t, err)

	s1 := GetSecret(&key1.PublicKey, key2)
	s2 := GetSecret(&key2.PublicKey, key1)

	assert.Equal(t, s1, s2)
}
