package beaconClient

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/attestantio/go-eth2-client/http"
)

type Client struct {
	*http.Service
}

func NewClient(url string) (newClient *Client, err error) {
	client, err := http.New(context.Background(),
		// WithAddress supplies the address of the beacon node, as a URL.
		http.WithLogLevel(zerolog.InfoLevel),
		http.WithAddress(url),
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
