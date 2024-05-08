package proofgen

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/pkg/errors"

	eigenpodproofs "github.com/Layr-Labs/eigenpod-proofs-generation"

	txsubmitter "github.com/Layr-Labs/eigenpod-proofs-generation/tx_submitoor/tx_submitter"

	"github.com/Layr-Labs/eigenpod-proofs-generation/beacon"

	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"

	commonutils "github.com/Layr-Labs/eigenpod-proofs-generation/common_utils"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func VerifyWithdrawalCredentialsGen(chainId uint64, provider *ethclient.Client, oracleStake, oracleHeader []byte, eigenPod common.Address, podOwner common.Address, validatorIndices []uint64) (*types.Transaction, error) {
	oracleBeaconBlockHeader, err := extractBlockHeader(oracleHeader)
	if err != nil {
		return nil, err
	}
	oracleStateJSON, err := parseDenebStateJSONFile(oracleStake)
	if err != nil {
		return nil, err
	}
	var oracleState deneb.BeaconState

	err = parseDenebBeaconStateFromJSON(*oracleStateJSON, &oracleState)
	if err != nil {
		return nil, err
	}

	versionedOracleState, err := beacon.CreateVersionedState(&oracleState)
	if err != nil {
		return nil, err
	}
	chainClient, err := txsubmitter.NewChainClient(provider, "", podOwner.Hex(), 0, 0)
	if err != nil {
		return nil, err
	}

	eigenPodProofs, err := eigenpodproofs.NewEigenPodProofs(chainId, 0)
	if err != nil {
		return nil, err
	}
	submitter := txsubmitter.NewEigenPodProofTxSubmitter(
		*chainClient,
		*eigenPodProofs,
	)
	return submitter.GenerateVerifyWithdrawalCredentialsTx(eigenPod, &versionedOracleState, &oracleBeaconBlockHeader, validatorIndices)
}

func extractBlockHeader(blockHeader []byte) (phase0.BeaconBlockHeader, error) {
	// Decode JSON
	var inputData commonutils.InputDataBlockHeader
	if err := json.Unmarshal(blockHeader, &inputData); err != nil {
		return phase0.BeaconBlockHeader{}, err
	}

	return inputData.Data.Header.Message, nil
}

func parseDenebStateJSONFile(oracleStake []byte) (*beaconStateJSONDeneb, error) {
	var beaconState beaconStateVersionDeneb
	err := json.Unmarshal(oracleStake, &beaconState)
	if err != nil {
		return nil, err
	}

	actualData := beaconState.Data
	return &actualData, nil
}

type beaconStateJSONDeneb struct {
	GenesisTime                  string                        `json:"genesis_time"`
	GenesisValidatorsRoot        string                        `json:"genesis_validators_root"`
	Slot                         string                        `json:"slot"`
	Fork                         *phase0.Fork                  `json:"fork"`
	LatestBlockHeader            *phase0.BeaconBlockHeader     `json:"latest_block_header"`
	BlockRoots                   []string                      `json:"block_roots"`
	StateRoots                   []string                      `json:"state_roots"`
	HistoricalRoots              []string                      `json:"historical_roots"`
	ETH1Data                     *phase0.ETH1Data              `json:"eth1_data"`
	ETH1DataVotes                []*phase0.ETH1Data            `json:"eth1_data_votes"`
	ETH1DepositIndex             string                        `json:"eth1_deposit_index"`
	Validators                   []*phase0.Validator           `json:"validators"`
	Balances                     []string                      `json:"balances"`
	RANDAOMixes                  []string                      `json:"randao_mixes"`
	Slashings                    []string                      `json:"slashings"`
	PreviousEpochParticipation   []string                      `json:"previous_epoch_participation"`
	CurrentEpochParticipation    []string                      `json:"current_epoch_participation"`
	JustificationBits            string                        `json:"justification_bits"`
	PreviousJustifiedCheckpoint  *phase0.Checkpoint            `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint   *phase0.Checkpoint            `json:"current_justified_checkpoint"`
	FinalizedCheckpoint          *phase0.Checkpoint            `json:"finalized_checkpoint"`
	InactivityScores             []string                      `json:"inactivity_scores"`
	CurrentSyncCommittee         *altair.SyncCommittee         `json:"current_sync_committee"`
	NextSyncCommittee            *altair.SyncCommittee         `json:"next_sync_committee"`
	LatestExecutionPayloadHeader *deneb.ExecutionPayloadHeader `json:"latest_execution_payload_header"`
	NextWithdrawalIndex          string                        `json:"next_withdrawal_index"`
	NextWithdrawalValidatorIndex string                        `json:"next_withdrawal_validator_index"`
	HistoricalSummaries          []*capella.HistoricalSummary  `json:"historical_summaries"`
}

type beaconStateJSONCapella struct {
	GenesisTime                  string                          `json:"genesis_time"`
	GenesisValidatorsRoot        string                          `json:"genesis_validators_root"`
	Slot                         string                          `json:"slot"`
	Fork                         *phase0.Fork                    `json:"fork"`
	LatestBlockHeader            *phase0.BeaconBlockHeader       `json:"latest_block_header"`
	BlockRoots                   []string                        `json:"block_roots"`
	StateRoots                   []string                        `json:"state_roots"`
	HistoricalRoots              []string                        `json:"historical_roots"`
	ETH1Data                     *phase0.ETH1Data                `json:"eth1_data"`
	ETH1DataVotes                []*phase0.ETH1Data              `json:"eth1_data_votes"`
	ETH1DepositIndex             string                          `json:"eth1_deposit_index"`
	Validators                   []*phase0.Validator             `json:"validators"`
	Balances                     []string                        `json:"balances"`
	RANDAOMixes                  []string                        `json:"randao_mixes"`
	Slashings                    []string                        `json:"slashings"`
	PreviousEpochParticipation   []string                        `json:"previous_epoch_participation"`
	CurrentEpochParticipation    []string                        `json:"current_epoch_participation"`
	JustificationBits            string                          `json:"justification_bits"`
	PreviousJustifiedCheckpoint  *phase0.Checkpoint              `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint   *phase0.Checkpoint              `json:"current_justified_checkpoint"`
	FinalizedCheckpoint          *phase0.Checkpoint              `json:"finalized_checkpoint"`
	InactivityScores             []string                        `json:"inactivity_scores"`
	CurrentSyncCommittee         *altair.SyncCommittee           `json:"current_sync_committee"`
	NextSyncCommittee            *altair.SyncCommittee           `json:"next_sync_committee"`
	LatestExecutionPayloadHeader *capella.ExecutionPayloadHeader `json:"latest_execution_payload_header"`
	NextWithdrawalIndex          string                          `json:"next_withdrawal_index"`
	NextWithdrawalValidatorIndex string                          `json:"next_withdrawal_validator_index"`
	HistoricalSummaries          []*capella.HistoricalSummary    `json:"historical_summaries"`
}

type beaconStateVersionDeneb struct {
	Data beaconStateJSONDeneb `json:"data"`
}

type beaconStateVersionCapella struct {
	Data beaconStateJSONCapella `json:"data"`
}

// nolint:gocyclo
func parseDenebBeaconStateFromJSON(data beaconStateJSONDeneb, s *deneb.BeaconState) error {
	var err error

	if data.GenesisTime == "" {
		return errors.New("genesis time missing")
	}
	if s.GenesisTime, err = strconv.ParseUint(data.GenesisTime, 10, 64); err != nil {
		return errors.Wrap(err, "invalid value for genesis time")
	}
	if data.GenesisValidatorsRoot == "" {
		return errors.New("genesis validators root missing")
	}
	genesisValidatorsRoot, err := hex.DecodeString(strings.TrimPrefix(data.GenesisValidatorsRoot, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid value for genesis validators root")
	}
	if len(genesisValidatorsRoot) != phase0.RootLength {
		return fmt.Errorf("incorrect length %d for genesis validators root", len(genesisValidatorsRoot))
	}
	copy(s.GenesisValidatorsRoot[:], genesisValidatorsRoot)
	if data.Slot == "" {
		return errors.New("slot missing")
	}
	slot, err := strconv.ParseUint(data.Slot, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for slot")
	}
	s.Slot = phase0.Slot(slot)
	if data.Fork == nil {
		return errors.New("fork missing")
	}
	s.Fork = data.Fork
	if data.LatestBlockHeader == nil {
		return errors.New("latest block header missing")
	}
	s.LatestBlockHeader = data.LatestBlockHeader
	if len(data.BlockRoots) == 0 {
		return errors.New("block roots missing")
	}
	s.BlockRoots = make([]phase0.Root, len(data.BlockRoots))
	for i := range data.BlockRoots {
		if data.BlockRoots[i] == "" {
			return fmt.Errorf("block root %d missing", i)
		}
		blockRoot, err := hex.DecodeString(strings.TrimPrefix(data.BlockRoots[i], "0x"))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid value for block root %d", i))
		}
		if len(blockRoot) != phase0.RootLength {
			return fmt.Errorf("incorrect length %d for block root %d", len(blockRoot), i)
		}
		copy(s.BlockRoots[i][:], blockRoot)
	}
	s.StateRoots = make([]phase0.Root, len(data.StateRoots))
	for i := range data.StateRoots {
		if data.StateRoots[i] == "" {
			return fmt.Errorf("state root %d missing", i)
		}
		stateRoot, err := hex.DecodeString(strings.TrimPrefix(data.StateRoots[i], "0x"))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid value for state root %d", i))
		}
		if len(stateRoot) != phase0.RootLength {
			return fmt.Errorf("incorrect length %d for state root %d", len(stateRoot), i)
		}
		copy(s.StateRoots[i][:], stateRoot)
	}
	s.HistoricalRoots = make([]phase0.Root, len(data.HistoricalRoots))
	for i := range data.HistoricalRoots {
		if data.HistoricalRoots[i] == "" {
			return fmt.Errorf("historical root %d missing", i)
		}
		historicalRoot, err := hex.DecodeString(strings.TrimPrefix(data.HistoricalRoots[i], "0x"))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid value for historical root %d", i))
		}
		if len(historicalRoot) != phase0.RootLength {
			return fmt.Errorf("incorrect length %d for historical root %d", len(historicalRoot), i)
		}
		copy(s.HistoricalRoots[i][:], historicalRoot)
	}
	if data.ETH1Data == nil {
		return errors.New("eth1 data missing")
	}
	s.ETH1Data = data.ETH1Data
	// ETH1DataVotes can be empty.
	s.ETH1DataVotes = data.ETH1DataVotes
	if data.Validators == nil {
		return errors.New("validators missing")
	}
	if data.ETH1DepositIndex == "" {
		return errors.New("eth1 deposit index missing")
	}
	if s.ETH1DepositIndex, err = strconv.ParseUint(data.ETH1DepositIndex, 10, 64); err != nil {
		return errors.Wrap(err, "invalid value for eth1 deposit index")
	}
	s.Validators = data.Validators
	s.Balances = make([]phase0.Gwei, len(data.Balances))
	for i := range data.Balances {
		if data.Balances[i] == "" {
			return fmt.Errorf("balance %d missing", i)
		}
		balance, err := strconv.ParseUint(data.Balances[i], 10, 64)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid value for balance %d", i))
		}
		s.Balances[i] = phase0.Gwei(balance)
	}
	s.RANDAOMixes = make([]phase0.Root, len(data.RANDAOMixes))
	for i := range data.RANDAOMixes {
		if data.RANDAOMixes[i] == "" {
			return fmt.Errorf("RANDAO mix %d missing", i)
		}
		randaoMix, err := hex.DecodeString(strings.TrimPrefix(data.RANDAOMixes[i], "0x"))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid value for RANDAO mix %d", i))
		}
		if len(randaoMix) != phase0.RootLength {
			return fmt.Errorf("incorrect length %d for RANDAO mix %d", len(randaoMix), i)
		}
		copy(s.RANDAOMixes[i][:], randaoMix)
	}
	s.Slashings = make([]phase0.Gwei, len(data.Slashings))
	for i := range data.Slashings {
		if data.Slashings[i] == "" {
			return fmt.Errorf("slashing %d missing", i)
		}
		slashings, err := strconv.ParseUint(data.Slashings[i], 10, 64)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid value for slashing %d", i))
		}
		s.Slashings[i] = phase0.Gwei(slashings)
	}
	s.PreviousEpochParticipation = make([]altair.ParticipationFlags, len(data.PreviousEpochParticipation))
	for i := range data.PreviousEpochParticipation {
		if data.PreviousEpochParticipation[i] == "" {
			return fmt.Errorf("previous epoch attestation %d missing", i)
		}
		previousEpochAttestation, err := strconv.ParseUint(data.PreviousEpochParticipation[i], 10, 8)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid value for previous epoch attestation %d", i))
		}
		s.PreviousEpochParticipation[i] = altair.ParticipationFlags(previousEpochAttestation)
	}
	s.CurrentEpochParticipation = make([]altair.ParticipationFlags, len(data.CurrentEpochParticipation))
	for i := range data.CurrentEpochParticipation {
		if data.CurrentEpochParticipation[i] == "" {
			return fmt.Errorf("current epoch attestation %d missing", i)
		}
		currentEpochAttestation, err := strconv.ParseUint(data.CurrentEpochParticipation[i], 10, 8)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid value for current epoch attestation %d", i))
		}
		s.CurrentEpochParticipation[i] = altair.ParticipationFlags(currentEpochAttestation)
	}
	if data.JustificationBits == "" {
		return errors.New("justification bits missing")
	}
	if s.JustificationBits, err = hex.DecodeString(strings.TrimPrefix(data.JustificationBits, "0x")); err != nil {
		return errors.Wrap(err, "invalid value for justification bits")
	}
	if data.PreviousJustifiedCheckpoint == nil {
		return errors.New("previous justified checkpoint missing")
	}
	s.PreviousJustifiedCheckpoint = data.PreviousJustifiedCheckpoint
	if data.CurrentJustifiedCheckpoint == nil {
		return errors.New("current justified checkpoint missing")
	}
	s.CurrentJustifiedCheckpoint = data.CurrentJustifiedCheckpoint
	if data.FinalizedCheckpoint == nil {
		return errors.New("finalized checkpoint missing")
	}
	s.FinalizedCheckpoint = data.FinalizedCheckpoint
	s.InactivityScores = make([]uint64, len(data.InactivityScores))
	for i := range data.InactivityScores {
		if data.InactivityScores[i] == "" {
			return fmt.Errorf("inactivity score %d missing", i)
		}
		if s.InactivityScores[i], err = strconv.ParseUint(data.InactivityScores[i], 10, 64); err != nil {
			return errors.Wrap(err, fmt.Sprintf("invalid value for inactivity score %d", i))
		}
	}
	if data.CurrentSyncCommittee == nil {
		return errors.New("current sync committee missing")
	}
	s.CurrentSyncCommittee = data.CurrentSyncCommittee
	if data.NextSyncCommittee == nil {
		return errors.New("next sync committee missing")
	}
	s.NextSyncCommittee = data.NextSyncCommittee
	s.LatestExecutionPayloadHeader = data.LatestExecutionPayloadHeader
	if data.NextWithdrawalIndex == "" {
		return errors.New("next withdrawal index missing")
	}
	nextWithdrawalIndex, err := strconv.ParseUint(data.NextWithdrawalIndex, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for next withdrawal index")
	}
	s.NextWithdrawalIndex = capella.WithdrawalIndex(nextWithdrawalIndex)
	if data.NextWithdrawalValidatorIndex == "" {
		return errors.New("next validator validator index missing")
	}
	nextWithdrawalValidatorIndex, err := strconv.ParseUint(data.NextWithdrawalValidatorIndex, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid value for next withdrawal validator index")
	}
	s.NextWithdrawalValidatorIndex = phase0.ValidatorIndex(nextWithdrawalValidatorIndex)
	if data.HistoricalSummaries == nil {
		return errors.New("historical summaries missing")
	}
	s.HistoricalSummaries = data.HistoricalSummaries

	return nil
}
