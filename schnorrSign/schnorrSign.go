package schnorrSign

import (
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
)

func GenerateKeyPair() ([]byte, []byte, error) {
	privKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, nil, err
	}
	pubKey := privKey.PubKey()
	return privKey.Serialize(), pubKey.SerializeCompressed(), nil
}

func Sign(privateKey, hash []byte) ([]byte, error) {
	pri, _ := btcec.PrivKeyFromBytes(privateKey)
	s, err := schnorr.Sign(pri, hash, schnorr.FastSign())
	if err != nil {
		return nil, err
	}
	return s.Serialize(), nil
}

func Verify(hash, sign, pubKey []byte) (bool, error) {
	publicKey, err := btcec.ParsePubKey(pubKey)
	if err != nil {
		return false, err
	}

	v, err := schnorr.ParseSignature(sign)
	if err != nil {
		return false, err
	}

	b := v.Verify(hash, publicKey)
	return b, nil
}
