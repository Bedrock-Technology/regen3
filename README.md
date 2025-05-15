# ClaimRewards

## Show Claimable Rewards

--earner-address is podOwner's address

```bash
~/go/bin/eigenlayer rewards show --eth-rpc-url https://mainnet.infura.io/v3/7a115a720b5d4bc1b2bddf39f9b89bfe --network mainnet --earner-address 0x3F4eaCeb930b0Edfa78a1DFCbaE5c5494E6e9850 --claim-type unclaimed --sidecar-http-rpc-url https://sidecar-rpc.eigenlayer.xyz/mainnet
```

## Generate ClaimRewards Proof

will generate proof in $PROOF_OUT_PATH/rewardProofs

```bash
./rewardProofs.sh
```

## BroadCast

```bash
./regen3 claimReward -c proconfig.yaml -p 0
```
