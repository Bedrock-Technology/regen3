package scanner

import (
	"github.com/Bedrock-Technology/regen3/beaconClient"
	"github.com/Bedrock-Technology/regen3/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type Scanner struct {
	Config       *config.Config
	DBEngine     *gorm.DB
	EthClient    *ethclient.Client
	BeaconClient *beaconClient.Client
	Quit         chan struct{}
}
