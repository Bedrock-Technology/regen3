package models

import (
	"gorm.io/gorm"
)

// Validator created during scan beacon chain, if a new deposit made to a pod address, CredentialVerified, WithdrawnOnChain,
// WithdrawnOnPod should be zero. if a pod need to do and did verifyWithdrawalCredentials, set CredentialVerified to 1.
// if a full withdrawal happened on chain. set WithdrawnOnChain to 1. WithdrawnOnPod set to 1 indicate that get pod event
// ValidatorWithdrawn and pod active minus 1.
// Number of CredentialVerified == 1 and WithdrawnOnPod == 0 should be equal to pod.active
type Validator struct {
	gorm.Model
	PubKey             string `gorm:"index:idx_pubKey,unique;type:varchar(128);not null;default:''" json:"pubKey"`
	PodAddress         string `gorm:"index:idx_validator,unique;type:varchar(128);not null;default:''" json:"podAddress"`
	ValidatorIndex     uint64 `gorm:"index:idx_validator,unique;not null;default:0" json:"validatorIndex"`
	CredentialVerified uint64 `gorm:"not null;default:0" json:"credentialVerified"`
	WithdrawnOnChain   uint64 `gorm:"not null;default:0" json:"withdrawnOnChain"`
	WithdrawnOnPod     uint64 `gorm:"not null;default:0" json:"withdrawnOnPod"`
	VoluntaryExit      uint64 `gorm:"not null;default:0" json:"voluntaryExit"`
}

// Pod for some pods will not enable checkPoint system, will not do verifyWithdrawalCredentials.
type Pod struct {
	gorm.Model
	Address      string `gorm:"index:idx_chain,unique;type:varchar(128);not null;default:''" json:"podAddress"`
	Owner        string `gorm:"index:idx_chain,unique;type:varchar(128);not null;default:''" json:"owner"`
	DelegateTo   string `gorm:"type:varchar(128);not null;default:''" json:"delegateTo"`
	PodIndex     uint64 `gorm:"not null;default:0" json:"podIndex"`
	IsCredential uint8  `gorm:"not null;default:0" json:"isCredential"`
}

// CheckPoint created during scan eth chain and get event of CheckpointCreated. ValidatorCheckpointed will change UnProofed,
// Proofed. CheckpointFinalized will set to 1 when CheckpointFinalized event emit.
type CheckPoint struct {
	gorm.Model
	Pod                 string `gorm:"index:idx_cp,unique;type:varchar(128);not null;default:''" json:"pod"`
	BeaconBlockRoot     string `gorm:"index:idx_cp,unique;type:varchar(128);not null;default:''" json:"beaconBlockRoot"`
	Proofed             string `gorm:"type:longblob;not null" json:"proofed"`
	Proofs              string `gorm:"type:longblob;not null" json:"proofs"`
	CheckpointTimestamp uint64 `gorm:"index:idx_cp,unique;not null;default:0" json:"checkpointTimestamp"`
	ActiveNum           uint64 `gorm:"not null;default:0" json:"activeNum"`
	BatchSize           uint64 `gorm:"not null;default:20" json:"BatchSize"`
	CheckpointFinalized uint64 `gorm:"not null;default:0" json:"checkpointFinalized"`
}

// QueueWithdrawals Completed set to 1 when get WithdrawalCompleted.
type QueueWithdrawals struct {
	gorm.Model
	Pod            string `gorm:"index:idx_qw,unique;type:varchar(128);not null;default:''" json:"pod"`
	WithdrawalRoot string `gorm:"index:idx_qw,unique;type:varchar(128);not null;default:''" json:"withdrawalRoot"`
	// json format
	Withdrawal string `gorm:"type:varchar(2048);not null;default:''" json:"withdrawal"`
	StartBlock uint64 `gorm:"not null;default:0" json:"startBlock"`
	Completed  uint64 `gorm:"not null;default:0" json:"completed"`
}

type Cursor struct {
	gorm.Model
	Meme string `gorm:"not null;default:''" json:"meme"`
	Slot uint64 `gorm:"not null;default:0" json:"slot"`
}

// Transaction record transaction fee that we spent
type Transaction struct {
	gorm.Model
	TxHash string `gorm:"type:varchar(128);not null;default:''" json:"txHash"`
	TxType string `gorm:"index:idx_qw;type:varchar(128);not null;default:''" json:"txType"`
	Fee    string `gorm:"type:varchar(128);not null;default:''" json:"fee"`
	Status uint64 `gorm:"not null;default:0" json:"status"`
	Block  uint64 `gorm:"not null;default:0" json:"block"`
}
