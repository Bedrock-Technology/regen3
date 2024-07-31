## go install github.com/ethereum/go-ethereum/cmd/abigen@latest
## abigen version 1.14.7-stable

abigen --abi ./contracts/DelegationManager/DelegationManager.json \
--pkg DelegationManager --out ./contracts/DelegationManager/DelegationManager.go

abigen --abi ./contracts/EigenPod/EigenPod.json \
--pkg EigenPod --out ./contracts/EigenPod/EigenPod.go

abigen --abi ./contracts/Restaking/Restaking.json \
--pkg Restaking --out ./contracts/Restaking/Restaking.go

abigen --abi ./contracts/Staking/Staking.json \
--pkg Staking --out ./contracts/Staking/Staking.go