package beaconClient

import (
	"context"
	"github.com/attestantio/go-eth2-client/api"
	"github.com/rs/zerolog"
	"time"

	"github.com/attestantio/go-eth2-client/http"
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
