## go install github.com/ethereum/go-ethereum/cmd/abigen@latest
## abigen version 1.14.0-stable

abigen --abi ./contracts/DelayedWithdrawalRouter/DelayedWithdrawalRouter.json \
--pkg DelayedWithdrawalRouter --out ./contracts/DelayedWithdrawalRouter/DelayedWithdrawalRouter.go

abigen --abi ./contracts/DelegationManager/DelegationManager.json \
--pkg DelegationManager --out ./contracts/DelegationManager/DelegationManager.go

abigen --abi ./contracts/EigenPod/EigenPod.json \
--pkg EigenPod --out ./contracts/EigenPod/EigenPod.go

abigen --abi ./contracts/EigenPodManager/EigenPodManager.json \
--pkg EigenPodManager --out ./contracts/EigenPodManager/EigenPodManager.go