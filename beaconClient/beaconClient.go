package beaconClient

import (
	"context"
	"encoding/hex"
	"strings"
	"time"

	"github.com/attestantio/go-eth2-client/api"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/rs/zerolog"
)

type Client struct {
	*http.Service
}

var (
	Cancel context.CancelFunc
	CTX    context.Context
)

func NewClient(url string) (newClient *Client, err error) {
	CTX, Cancel = context.WithCancel(context.Background())
	client, err := http.New(CTX,
		// WithAddress supplies the address of the beacon node, as a URL.
		http.WithLogLevel(zerolog.InfoLevel),
		http.WithAddress(url),
		http.WithTimeout(time.Second*60),
	)
	if err != nil {
		return nil, err
	}
	service := client.(*http.Service)
	newClient = &Client{
		service,
	}
	return
}

// GetLatestSlotNumber minus 2 slot for security
func (c *Client) GetLatestSlotNumber() (slotNumber uint64, err error) {
	header, err := c.BeaconBlockHeader(CTX, &api.BeaconBlockHeaderOpts{
		Block: "head",
	})
	if err != nil {
		return 0, err
	}
	return uint64(header.Data.Header.Message.Slot - 2), nil
}

func mustParsePubKey(input string) *phase0.BLSPubKey {
	pubKey, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic("invalid public key")
	}
	if len(pubKey) != phase0.PublicKeyLength {
		panic("invalid length public key")
	}

	var res phase0.BLSPubKey
	copy(res[:], pubKey)

	return &res
}

func (c *Client) ValidatorsByPubkeys(pubKeys []string) (*api.Response[map[phase0.ValidatorIndex]*apiv1.Validator], error) {
	blsPubKeys := []phase0.BLSPubKey{}
	for _, v := range pubKeys {
		blsPubKeys = append(blsPubKeys, *mustParsePubKey(v))
	}
	return c.Validators(CTX, &api.ValidatorsOpts{
		State:   "head",
		PubKeys: blsPubKeys,
	})
}

func (c *Client) ValidatorsByIndex(index []uint64) (*api.Response[map[phase0.ValidatorIndex]*apiv1.Validator], error) {
	validatorIndex := []phase0.ValidatorIndex{}
	for _, v := range index {
		validatorIndex = append(validatorIndex, phase0.ValidatorIndex(v))
	}
	return c.Validators(CTX, &api.ValidatorsOpts{
		State:   "head",
		Indices: validatorIndex,
	})
}
