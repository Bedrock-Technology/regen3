package schnorrSign

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	privatKey, pubKey, err := GenerateKeyPair()
	if err != nil {
		t.Log("err:", err)
		return
	}
	t.Logf("priv[%s],pub[%s]", hex.EncodeToString(privatKey), hex.EncodeToString(pubKey))
}

func TestSign(t *testing.T) {
	privatKey, pubKey, err := GenerateKeyPair()
	if err != nil {
		t.Log("err:", err)
		return
	}
	t.Logf("priv[%s],pub[%s]", hex.EncodeToString(privatKey), hex.EncodeToString(pubKey))

	pri, pub := btcec.PrivKeyFromBytes(privatKey)
	s, err := schnorr.Sign(pri, []byte(`adddddddddsdfsadfdsafdsafdsafdsa`), schnorr.FastSign())
	if err != nil {
		t.Log("err:", err)
		return
	}
	t.Log("sign:", hex.EncodeToString(s.Serialize()))

	//v
	v, err := schnorr.ParseSignature(s.Serialize())
	if err != nil {
		t.Log("err:", err)
		return
	}

	b := v.Verify([]byte(`adddddddddsdfsadfdsafdsafdsafdsa`), pub)
	t.Log("verify:", b)
}

//replace hash to sign cause sign already contains hash
