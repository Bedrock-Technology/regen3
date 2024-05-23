package keyAgentRpc

type Request struct {
	Content string `json:"content"`
	Sign    string `json:"sign"`
	PubKey  string `json:"pubKey"`
}

type Response struct {
	Content string `json:"content"`
	Sign    string `json:"sign"`
}

type Eth1AddressRequest struct {
	Service uint64 `json:"service"`
	Index   uint64 `json:"index"`
}

type Eth1AddressResponse struct {
	Index   uint64 `json:"index"`
	Service uint64 `json:"service"`
	Address string `json:"address"`
}

type Eth1SignRequest struct {
	Service    uint64 `json:"service"`
	Index      uint64 `json:"index"`
	DigestHash string `json:"digestHash"`
}

type Eth1SignResponse struct {
	Index   uint64 `json:"index"`
	Service uint64 `json:"service"`
	Sig     string `json:"sig"`
}
