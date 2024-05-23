package keyAgentRpc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Bedrock-Technology/regen3/config"
	"github.com/Bedrock-Technology/regen3/ecdh"
	"github.com/Bedrock-Technology/regen3/schnorrSign"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"net/http"
)

type Client struct {
	endpoint string
	imp      *http.Client
	config   config.KeyAgent
}

func NewClient(conf config.KeyAgent) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}

	return &Client{
		imp:      client,
		endpoint: conf.KeyAgentRpc,
		config:   conf,
	}
}

func successCode(code int) bool {
	codes := []int{http.StatusOK, http.StatusCreated}
	for _, v := range codes {
		if code == v {
			return true
		}
	}
	return false
}

func checkResponse(rsp *http.Response) (body []byte, err error) {
	defer rsp.Body.Close()

	body, err = io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	if !successCode(rsp.StatusCode) {
		err = errors.New(fmt.Sprintf("Http Error Status Code:%v,Body:%v", rsp.StatusCode, string(body)))
		return nil, err
	}

	return body, nil
}

func (client *Client) Post(url string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return client.imp.Do(req)
}

func (client *Client) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	return client.imp.Do(req)
}

func (client *Client) Eth1SignTransaction(digestHash []byte) (sig []byte, err error) {
	signPri, _ := hex.DecodeString(client.config.SignPri)
	signPub, _ := hex.DecodeString(client.config.KeyAgentSignPub)
	keyAgentEncryptPub, _ := hex.DecodeString(client.config.KeyAgentEncryptPub)
	keyAgentEncryptPubKey, _ := crypto.UnmarshalPubkey(keyAgentEncryptPub)
	encryptPri, _ := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	secret := ecdh.GetSecret(keyAgentEncryptPubKey, encryptPri)
	aesKey := pbkdf2.Key(secret.Bytes(), []byte(`Rockx`), 1024, 16, sha1.New)
	//
	eth1SignReq := Eth1SignRequest{
		Service:    client.config.Service,
		Index:      client.config.Index,
		DigestHash: hex.EncodeToString(digestHash),
	}
	content, _ := json.Marshal(&eth1SignReq)
	contentEncypt, _ := AesEncrypt(content, aesKey)

	sign, _ := Sign(contentEncypt, signPri)
	eth1req := Request{
		Content: hex.EncodeToString(contentEncypt),
		Sign:    hex.EncodeToString(sign),
		PubKey:  hex.EncodeToString(crypto.FromECDSAPub(&encryptPri.PublicKey)),
	}
	req, _ := json.Marshal(&eth1req)
	resp, err := client.Post(fmt.Sprintf("%s/eth1/sign", client.endpoint), bytes.NewReader(req))
	if err != nil {
		return
	}
	body, err := checkResponse(resp)
	if err != nil {
		return
	}
	respEth1 := Response{}
	_ = json.Unmarshal(body, &respEth1)
	respSign, _ := hex.DecodeString(respEth1.Sign)
	respContent, _ := hex.DecodeString(respEth1.Content)
	if !Verify(respContent, respSign, [][]byte{signPub}) {
		err = fmt.Errorf("verify error")
		return
	}
	//decode
	respContentDecrypt, _ := AesDecrypt(respContent, aesKey)
	eth1SignResponse := Eth1SignResponse{}
	_ = json.Unmarshal(respContentDecrypt, &eth1SignResponse)
	sig, _ = hex.DecodeString(eth1SignResponse.Sig)

	return
}

func (client *Client) Eth1Address() (address string, err error) {
	signPri, _ := hex.DecodeString(client.config.SignPri)
	signPub, _ := hex.DecodeString(client.config.KeyAgentSignPub)
	keyAgentEncryptPub, _ := hex.DecodeString(client.config.KeyAgentEncryptPub)
	keyAgentEncryptPubKey, _ := crypto.UnmarshalPubkey(keyAgentEncryptPub)
	encryptPri, _ := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	secret := ecdh.GetSecret(keyAgentEncryptPubKey, encryptPri)
	aesKey := pbkdf2.Key(secret.Bytes(), []byte(`Rockx`), 1024, 16, sha1.New)
	//
	eth1AddressReq := Eth1AddressRequest{
		Service: client.config.Service,
		Index:   client.config.Index,
	}
	content, _ := json.Marshal(&eth1AddressReq)
	contentEncypt, _ := AesEncrypt(content, aesKey)

	sign, _ := Sign(contentEncypt, signPri)
	eth1req := Request{
		Content: hex.EncodeToString(contentEncypt),
		Sign:    hex.EncodeToString(sign),
		PubKey:  hex.EncodeToString(crypto.FromECDSAPub(&encryptPri.PublicKey)),
	}
	req, _ := json.Marshal(&eth1req)
	resp, err := client.Post(fmt.Sprintf("%s/eth1/address", client.endpoint), bytes.NewReader(req))
	if err != nil {
		return
	}
	body, err := checkResponse(resp)
	if err != nil {
		return
	}
	respEth1 := Response{}
	_ = json.Unmarshal(body, &respEth1)
	respSign, _ := hex.DecodeString(respEth1.Sign)
	respContent, _ := hex.DecodeString(respEth1.Content)
	if !Verify(respContent, respSign, [][]byte{signPub}) {
		err = fmt.Errorf("verify error")
		return
	}
	//decode
	respContentDecrypt, _ := AesDecrypt(respContent, aesKey)
	eth1AddressResponse := Eth1AddressResponse{}
	_ = json.Unmarshal(respContentDecrypt, &eth1AddressResponse)
	address = eth1AddressResponse.Address

	return
}

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pKCS7UnPadding(origData)
	return origData, nil
}

func Sign(data, privateKey []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	sign, err := schnorrSign.Sign(privateKey, hash[:])
	if err != nil {
		return nil, err
	}

	return sign, nil
}

func Verify(data, sign []byte, publicKeys [][]byte) bool {
	hash := sha256.Sum256(data)
	verify := false
	for i := len(publicKeys) - 1; i >= 0; i-- {
		isVerified, err := schnorrSign.Verify(hash[:], sign, publicKeys[i])
		if err != nil {
		}
		if isVerified {
			verify = true
			break
		} else {
		}
	}
	if !verify {
		return false
	}

	return true
}
