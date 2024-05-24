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
	PubKey         string `gorm:"index:idx_pubKey,unique;type:varchar(128);not null;default:''" json:"pubKey"`
	ValidatorIndex uint64 `gorm:"index:idx_validator,unique;not null;default:0" json:"validatorIndex"`
	PodAddress     string `gorm:"index:idx_validator,unique;type:varchar(128);not null;default:''" json:"podAddress"`
	// Eth Block Number
	CredentialVerified uint64 `gorm:"not null;default:0" json:"credentialVerified"`
	// Beacon Slot
	WithdrawnOnChain uint64 `gorm:"not null;default:0" json:"withdrawnOnChain"`
	// Eth Block Number
	WithdrawnOnPod uint64 `gorm:"not null;default:0" json:"withdrawnOnPod"`
	// Beacon Slot
	VoluntaryExit uint64 `gorm:"not null;default:0" json:"voluntaryExit"`
}

// Pod for some pods will not enable checkPoint system, will not do verifyWithdrawalCredentials.
type Pod struct {
	gorm.Model
	Address      string `gorm:"index:idx_chain,unique;type:varchar(128);not null;default:''" json:"podAddress"`
	Owner        string `gorm:"index:idx_chain,unique;type:varchar(128);not null;default:''" json:"owner"`
	PodIndex     uint64 `gorm:"not null;default:0" json:"podIndex"`
	DelegateTo   string `gorm:"type:varchar(128);not null;default:''" json:"delegateTo"`
	IsCredential uint8  `gorm:"not null;default:0" json:"isCredential"`
}

// CheckPoint created during scan eth chain and get event of CheckpointCreated. ValidatorCheckpointed will change UnProofed,
// Proofed. CheckpointFinalized will set to 1 when CheckpointFinalized event emit.
type CheckPoint struct {
	gorm.Model
	Pod                 string `gorm:"index:idx_cp,unique;type:varchar(128);not null;default:''" json:"pod"`
	CheckpointTimestamp uint64 `gorm:"index:idx_cp,unique;not null;default:0" json:"checkpointTimestamp"`
	BeaconBlockRoot     string `gorm:"index:idx_cp,unique;type:varchar(128);not null;default:''" json:"beaconBlockRoot"`
	//json array format: [123,45456,validatorIndex]
	Proofed string `gorm:"type:blob;not null" json:"proofed"`
	//json array format: [123,45456,validatorIndex]
	UnProofed string `gorm:"type:blob;not null" json:"unProofed"`
	// check activeNum form Validator equal to pod.active on the block
	ActiveNum uint64 `gorm:"not null;default:0" json:"activeNum"`
	// if CheckpointFinalized set to 1, check len(Proofed) equal to ActiveNum
	CheckpointFinalized uint8 `gorm:"not null;default:0" json:"checkpointFinalized"`
}

// QueueWithdrawals Completed set to 1 when get WithdrawalCompleted.
type QueueWithdrawals struct {
	gorm.Model
	Pod            string `gorm:"index:idx_qw,unique;type:varchar(128);not null;default:''" json:"pod"`
	WithdrawalRoot string `gorm:"index:idx_qw,unique;type:varchar(128);not null;default:''" json:"withdrawalRoot"`
	//json format
	Withdrawal string `gorm:"type:varchar(2048);not null;default:''" json:"withdrawal"`
	StartBlock uint64 `gorm:"not null;default:0" json:"startBlock"`
	Completed  uint8  `gorm:"not null;default:0" json:"completed"`
}

type Cursor struct {
	gorm.Model
	//slot start to process
	Slot uint64 `gorm:"not null;default:0" json:"slot"`
	Meme string `gorm:"not null;default:''" json:"meme"`
}

// Transaction record transaction fee that we spent
type Transaction struct {
	gorm.Model
	TxHash string `gorm:"type:varchar(128);not null;default:''" json:"txHash"`
	Status uint64 `gorm:"not null;default:0" json:"status"`
	TxType string `gorm:"type:varchar(128);not null;default:''" json:"txType"`
	Fee    string `gorm:"type:varchar(128);not null;default:''" json:"fee"`
}
