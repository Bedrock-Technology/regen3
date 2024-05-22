package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Addresses struct {
	stakingContract                string
	restakingContract              string
	eigenDelegationManagerContract string
	eigenOracleContract            string
}

var MainnetAddresses = Addresses{
	stakingContract:                "",
	restakingContract:              "",
	eigenDelegationManagerContract: "",
	eigenOracleContract:            "",
}

var HoleskyAddresses = Addresses{
	stakingContract:                "0x0f59BfDEdbB4ECc965be28484BfD968552fD5C67",
	restakingContract:              "0xf59fc684Ad69A7F8B1C563D8b9fC4003F841F4Ef",
	eigenDelegationManagerContract: "0xA44151489861Fe9e3055d95adC98FbD462B948e7",
	eigenOracleContract:            "0x4C116BB629bff7A8373c2378bBd919f8349B8f25",
}

type Config struct {
	Network                        string    `yaml:"network"`
	ChainId                        uint64    `yaml:"-"`
	EthClient                      string    `yaml:"ethClient"`
	BeaconClient                   string    `yaml:"beaconClient"`
	MysqlDsn                       string    `yaml:"mysqlDsn"`
	LogLevel                       string    `yaml:"logLevel"`
	SlackUrl                       string    `yaml:"slackUrl"`
	DataPath                       string    `yaml:"dataPath"`
	StakingContract                string    `yaml:"-"`
	RestakingContract              string    `yaml:"-"`
	EigenDelegationManagerContract string    `yaml:"-"`
	EigenOracleContract            string    `yaml:"-"`
	CheckVerifyWithdrawCredential  TimerSpec `yaml:"checkVerifyWithdrawCredential"`
}

type TimerSpec struct {
	IntervalBlock uint64 `yaml:"intervalBlock"`
	FirstRun      uint64 `yaml:"firstRun"`
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
	if config.Network == "holesky" {
		config.StakingContract = HoleskyAddresses.stakingContract
		config.RestakingContract = HoleskyAddresses.restakingContract
		config.EigenDelegationManagerContract = HoleskyAddresses.eigenDelegationManagerContract
		config.EigenOracleContract = HoleskyAddresses.eigenOracleContract
		config.ChainId = 17000
	} else if config.Network == "mainnet" {
		config.StakingContract = MainnetAddresses.stakingContract
		config.RestakingContract = MainnetAddresses.restakingContract
		config.EigenDelegationManagerContract = MainnetAddresses.eigenDelegationManagerContract
		config.EigenOracleContract = MainnetAddresses.eigenOracleContract
		config.ChainId = 1
	} else {
		panic("invalid network")
	}
	return
}
