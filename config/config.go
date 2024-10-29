package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Addresses struct {
	stakingContract                string
	restakingContract              string
	eigenDelegationManagerContract string
}

var MainnetAddresses = Addresses{
	stakingContract:                "0x4beFa2aA9c305238AA3E0b5D17eB20C045269E9d",
	restakingContract:              "0x3F4eaCeb930b0Edfa78a1DFCbaE5c5494E6e9850",
	eigenDelegationManagerContract: "0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A",
}

var HoleskyAddresses = Addresses{
	stakingContract:                "0x0f59BfDEdbB4ECc965be28484BfD968552fD5C67",
	restakingContract:              "0xf59fc684Ad69A7F8B1C563D8b9fC4003F841F4Ef",
	eigenDelegationManagerContract: "0xA44151489861Fe9e3055d95adC98FbD462B948e7",
}

const (
	HoleskyMinWithdrawalDelayBlocks = 3600 // half day
	MainnetMinWithdrawalDelayBlocks = 50403
)

type Config struct {
	ReportSpec                     string    `yaml:"reportSpec"`
	MysqlDsn                       string    `yaml:"mysqlDsn"`
	EthClient                      string    `yaml:"ethClient"`
	Network                        string    `yaml:"network"`
	BeaconClient                   string    `yaml:"beaconClient"`
	EigenDelegationManagerContract string    `yaml:"-"`
	LogLevel                       string    `yaml:"logLevel"`
	SlackUrl                       string    `yaml:"slackUrl"`
	RestakingContract              string    `yaml:"-"`
	StakingContract                string    `yaml:"-"`
	KeyAgent                       KeyAgent  `yaml:"keyAgent"`
	CheckVerifyWithdrawCredential  TimerSpec `yaml:"checkVerifyWithdrawCredential"`
	CheckStartCheckPoint           TimerSpec `yaml:"checkStartCheckPoint"`
	CheckQueueWithdraw             TimerSpec `yaml:"checkQueueWithdraw"`
	CheckVerifyCheckPoint          TimerSpec `yaml:"checkVerifyCheckPoint"`
	ChainId                        uint64    `yaml:"-"`
	CheckPointThreshold            uint64    `yaml:"-"`
	Pod0CheckPointThreshold        uint64    `yaml:"-"`
	MinWithdrawalDelayBlocks       uint64    `yaml:"-"`
}

type KeyAgent struct {
	KeyAgentRpc        string `yaml:"keyAgentRpc"`
	Address            string `yaml:"address"`
	SignPri            string `yaml:"signPri"`
	KeyAgentSignPub    string `yaml:"keyAgentSignPub"`
	KeyAgentEncryptPub string `yaml:"keyAgentEncryptPub"`
	Service            uint64 `yaml:"service"`
	Index              uint64 `yaml:"index"`
}

type TimerSpec struct {
	Enable        bool   `yaml:"enable"`
	IntervalBlock uint64 `yaml:"intervalBlock"`
	FirstRun      uint64 `yaml:"firstRun"`
	BatchSize     int    `yaml:"batchSize"`
}

func LoadConfig(path string) (config *Config) {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}
	switch config.Network {
	case "holesky":
		config.StakingContract = HoleskyAddresses.stakingContract
		config.RestakingContract = HoleskyAddresses.restakingContract
		config.EigenDelegationManagerContract = HoleskyAddresses.eigenDelegationManagerContract
		config.MinWithdrawalDelayBlocks = HoleskyMinWithdrawalDelayBlocks
		config.ChainId = 17000
		config.CheckPointThreshold = 1e9
		config.Pod0CheckPointThreshold = 1e9
	case "mainnet":
		config.StakingContract = MainnetAddresses.stakingContract
		config.RestakingContract = MainnetAddresses.restakingContract
		config.EigenDelegationManagerContract = MainnetAddresses.eigenDelegationManagerContract
		config.MinWithdrawalDelayBlocks = MainnetMinWithdrawalDelayBlocks
		config.ChainId = 1
		config.CheckPointThreshold = 5 * 32e9
		config.Pod0CheckPointThreshold = 20 * 32e9
	default:
		panic("invalid network")
	}
	return
}
