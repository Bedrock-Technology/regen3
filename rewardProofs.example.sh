#!/bin/bash
#go install github.com/Layr-Labs/eigenlayer-cli/cmd/eigenlayer@latest
#./eigenlayer rewards claim --eth-rpc-url --network --earner-address --output-type json --output-file

ETH_RPC_URL=""
NETWORK=""
EIGEN_PATH="/home/ubuntu/go/bin"
PROOF_OUT_PATH="."
#pods array, using space to separate.
PODOWNERS=()

for ((index = 0; index < ${#PODOWNERS[@]}; index++)); do
    #echo ${array[$index]}
    pod="pod_$index"
    owner=${PODOWNERS[$index]}
    echo "generate $pod, owner:$owner"
    $EIGEN_PATH/eigenlayer rewards claim --eth-rpc-url "$ETH_RPC_URL" --network "$NETWORK" --earner-address "$owner" --output-type json --output-file $PROOF_OUT_PATH/rewardProofs/${pod}.json --silent
done
