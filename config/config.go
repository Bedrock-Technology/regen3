package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Addresses struct {
	stakingContract                string
	stakingPectraContract          string
	restakingContract              string
	restakingPectraContract        string
	eigenDelegationManagerContract string
	rewardCoordinator              string
	eigenToken                     string
	eigenPodManager                string
}

var MainnetAddresses = Addresses{
	stakingContract:                "0x4beFa2aA9c305238AA3E0b5D17eB20C045269E9d",
	restakingContract:              "0x3F4eaCeb930b0Edfa78a1DFCbaE5c5494E6e9850",
	eigenDelegationManagerContract: "0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A",
	rewardCoordinator:              "0x7750d328b314EfFa365A0402CcfD489B80B0adda",
	eigenToken:                     "0xec53bF9167f50cDEB3Ae105f56099aaaB9061F83",
	eigenPodManager:                "0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338	",
}

var HoleskyAddresses = Addresses{
	stakingContract:                "0x0f59BfDEdbB4ECc965be28484BfD968552fD5C67",
	restakingContract:              "0xf59fc684Ad69A7F8B1C563D8b9fC4003F841F4Ef",
	eigenDelegationManagerContract: "0xA44151489861Fe9e3055d95adC98FbD462B948e7",
	rewardCoordinator:              "0xAcc1fb458a1317E886dB376Fc8141540537E68fE",
	eigenToken:                     "0xdeeeeE2b48C121e6728ed95c860e296177849932",
	eigenPodManager:                "0x30770d7E3e71112d7A6b7259542D1f680a70e315",
}

var HoodiAddresses = Addresses{
	stakingContract:                "0xd8B81B8950981EFbA4c00Eed567f903580A6649c",
	restakingContract:              "0xdF1925B7A0f56a3ED7f74bE2a813Ae8bbA756e59",
	eigenDelegationManagerContract: "0x867837a9722C512e0862d8c2E15b8bE220E8b87d",
	rewardCoordinator:              "0x29e8572678e0c272350aa0b4B8f304E47EBcd5e7",
	eigenToken:                     "0x8ae2520954db7D80D66835cB71E692835bbA45bf",
	stakingPectraContract:          "0x83ED17AAe050335E3d459EF7867672f166d25995",
	restakingPectraContract:        "0x4940eE4f0Ff6dAb57Db44Cd71683Aab0ae9cf2c4",
	eigenPodManager:                "0xcd1442415Fc5C29Aa848A49d2e232720BE07976c",
}

const (
	HoleskyMinWithdrawalDelayBlocks = 3600 // half day
	HoodiMinWithdrawalDelayBlocks   = 3600 // half day
	MainnetMinWithdrawalDelayBlocks = 100803
)

type Config struct {
	ReportSpec                     string    `yaml:"reportSpec"`
	MysqlDsn                       string    `yaml:"mysqlDsn"`
	EthClient                      string    `yaml:"ethClient"`
	Network                        string    `yaml:"network"`
	BeaconClient                   string    `yaml:"beaconClient"`
	EigenDelegationManagerContract string    `yaml:"-"`
	EigenToken                     string    `yaml:"-"`
	LogLevel                       string    `yaml:"logLevel"`
	SlackUrl                       string    `yaml:"slackUrl"`
	RestakingContract              string    `yaml:"-"`
	StakingContract                string    `yaml:"-"`
	RestakingPectraContract        string    `yaml:"-"`
	StakingPectraContract          string    `yaml:"-"`
	EigenPodManager                string    `yaml:"-"`
	KeyAgent                       KeyAgent  `yaml:"keyAgent"`
	CheckVerifyWithdrawCredential  TimerSpec `yaml:"checkVerifyWithdrawCredential"`
	CheckStartCheckPoint           TimerSpec `yaml:"checkStartCheckPoint"`
	CheckQueueWithdraw             TimerSpec `yaml:"checkQueueWithdraw"`
	CheckVerifyCheckPoint          TimerSpec `yaml:"checkVerifyCheckPoint"`
	ChainId                        uint64    `yaml:"-"`
	CheckPointThreshold            uint64    `yaml:"-"`
	Pod0CheckPointThreshold        uint64    `yaml:"-"`
	MinWithdrawalDelayBlocks       uint64    `yaml:"-"`
	RewardCoordinator              string    `yaml:"-"`
	// In Gwei
	EigenTokenThreshold uint64 `yaml:"-"`
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
		config.RewardCoordinator = HoleskyAddresses.rewardCoordinator
		config.EigenPodManager = HoleskyAddresses.eigenPodManager
		config.EigenToken = HoleskyAddresses.eigenToken
		config.MinWithdrawalDelayBlocks = HoleskyMinWithdrawalDelayBlocks
		config.ChainId = 17000
		config.CheckPointThreshold = 1e9
		config.Pod0CheckPointThreshold = 1e9
		config.EigenTokenThreshold = 1_000 * 1e9
	case "mainnet":
		config.StakingContract = MainnetAddresses.stakingContract
		config.RestakingContract = MainnetAddresses.restakingContract
		config.EigenPodManager = MainnetAddresses.eigenPodManager
		config.EigenDelegationManagerContract = MainnetAddresses.eigenDelegationManagerContract
		config.RewardCoordinator = MainnetAddresses.rewardCoordinator
		config.EigenToken = MainnetAddresses.eigenToken
		config.MinWithdrawalDelayBlocks = MainnetMinWithdrawalDelayBlocks
		config.ChainId = 1
		config.CheckPointThreshold = 1 * 32e9
		config.Pod0CheckPointThreshold = 20 * 32e9
		config.EigenTokenThreshold = 500 * 1e9
	case "hoodi":
		config.StakingContract = HoodiAddresses.stakingContract
		config.StakingPectraContract = HoleskyAddresses.stakingPectraContract
		config.RestakingContract = HoodiAddresses.restakingContract
		config.RestakingPectraContract = HoodiAddresses.restakingPectraContract
		config.EigenPodManager = HoodiAddresses.eigenPodManager
		config.EigenDelegationManagerContract = HoodiAddresses.eigenDelegationManagerContract
		config.RewardCoordinator = HoodiAddresses.rewardCoordinator
		config.EigenToken = HoodiAddresses.eigenToken
		config.MinWithdrawalDelayBlocks = HoodiMinWithdrawalDelayBlocks
		config.ChainId = 560048
		config.CheckPointThreshold = 1e9
		config.Pod0CheckPointThreshold = 1e9
		config.EigenTokenThreshold = 1_000 * 1e9
	default:
		panic("invalid network")
	}
	return
}
