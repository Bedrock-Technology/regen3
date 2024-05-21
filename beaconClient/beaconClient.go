package beaconClient

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/attestantio/go-eth2-client/http"
)

func NewClient(url string) (service *http.Service, err error) {
	client, err := http.New(context.Background(),
		// WithAddress supplies the address of the beacon node, as a URL.
		http.WithLogLevel(zerolog.InfoLevel),
		http.WithAddress(url),
	)
	if err != nil {
		return nil, err
	}
	service = client.(*http.Service)
	return
}
