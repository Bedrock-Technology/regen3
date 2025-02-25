package EigenPod

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/Bedrock-Technology/regen3/contracts/Restaking"
	"github.com/ethereum/go-ethereum"

	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var provider *ethclient.Client
var contract *EigenPod

var (
	RpcHost = ""
	ChainId = int64(0)
)

var (
	ContractAddress = ""
	PodOwner        = ""
	Pod             = ""
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	RpcHost = os.Getenv("HOLESKY_URL")
	ChainId, _ = strconv.ParseInt(os.Getenv("HOLESKY_CHAIN_ID"), 0, 64)
	ContractAddress = os.Getenv("HOLESKY_POD")
	PodOwner = os.Getenv("HOLESKY_POD_WONER")
	Pod = os.Getenv("HOLESKY_POD")

	client, err := ethclient.Dial(RpcHost)
	if err != nil {
		panic(err)
	}
	d, err := NewEigenPod(common.HexToAddress(ContractAddress), client)
	if err != nil {
		panic(err)
	}
	provider = client
	contract = d
}

type ValidatorStatus uint8

func (vs ValidatorStatus) String() string {
	switch vs {
	case ValidatorInactive:
		return "INACTIVE"
	case ValidatorActive:
		return "ACTIVE"
	case ValidatorWithdrawn:
		return "WITHDRAWN"
	}
	return "UNKNOWN"
}

const (
	ValidatorInactive  ValidatorStatus = 0
	ValidatorActive    ValidatorStatus = 1
	ValidatorWithdrawn ValidatorStatus = 2
)

func TestEigenPodCaller_ValidatorPubkeyToInfo(t *testing.T) {
	pubKey := "aee4bc69d8ea4997cea56fb7aa8a56af42a4f8c41105651b6deb316ebc436a5b43f7c6657401fdee88b2e891d0cb42a3"
	pubKeyByte, _ := hex.DecodeString(pubKey)
	validatorInfo, err := contract.ValidatorPubkeyToInfo(&bind.CallOpts{}, pubKeyByte)
	if err != nil {
		t.Fatalf("failed to call ValidatorPubkeyToInfo: %v", err)
	}
	t.Logf("ValidatorPubkeyToInfo: %v", validatorInfo)
	if ValidatorStatus(validatorInfo.Status) == ValidatorActive {
		t.Logf("ValidatorPubkeyToInfo status: %v", "Active")
	}
	t.Logf("ValidatorPubkeyToInfo Status: %v", ValidatorStatus(validatorInfo.Status))
}

func TestFilter(t *testing.T) {
	logs, err := provider.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(1362421)),
		ToBlock:   big.NewInt(int64(1362421)),
		Addresses: []common.Address{common.HexToAddress(Pod)},
	})
	if err != nil {
		t.Logf("err:%v", err)
		return
	}
	out, _ := json.Marshal(&logs)
	t.Logf("logs:%v", string(out))
}

func TestDecode(t *testing.T) {
	data := "3f65cf19000000000000000000000000000000000000000000000000000000006653e86000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000007c0693b21bd86e79438b2872fb7203a95ec456e46b1c3349b783d84637a1769077a000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000608a049ca5e244b93927c28bfe45522b6d96a4e5c6f20567b030b78bed03f5123e5497978b5956205700cb0904c604d7a0d777d77f957c03960c3fcdb052335bf9d781fce6ab374318a8e6b8e812178d230a990a7b4d99a6d42e8056a3dc887986000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000001a1c1e0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000005c0bda5ad5d4261ff9982a9c82102f72454b294d3668aa6873dcd916fbab308b57ed919dddfbaef5c613efb89f58374f0c7f646d21c7ee6b01258fb99bd2df3e8bad5a2e5cafd2b1584313237fb946bbba8e2b8777f04ad2fb7366d26761971e9d87ec495d66ceb2e9afeed45b112f1b6fc7edf39572434f1fb6b620a315431dc3e79520eca463e0bc1dc7f8b6e25b2aeecf95dc94deeb53b8d2f3065599a8ab968ac53dc945ecc755b8a8306f66f28cbd68eb7d0a6641a65fec51964064b773e7d0c2bf61b2b9dd9d556655749169f9f19998cc79fcdceca0483ac44405991e66a76a85f4db1cca79a6eddf8e58d8bd1f42c99c1c99385ccafc3128fdf99b6e1f1e5dd7b3e94a3627828c3c4d34c8ea44ec2e5a9fc943c023e861a3b10fe1b67c87b48ec15f8d8c64574899e31b30e4987931f7e46ec04d908d7db3fcd4935d16a6c952c610c86d53ebed5539f3070ae8770b062f6b97f9317d47877a3d9efcc0b88b2a9c8687497122d1ec6bce6d9a052c193e85c79d16dc9d03596f05dc1188358a7ca2b8279fd0b7f416c389cbc3c2f0d83f39d68e26dc0fff4d296d3a9834d23fc8e359bd64a5f117d3c5b7ee327c88596f098e5e116e89236d18feb4b726ab58d900f5e182e3c50ef74969ea16c7726c549757cc23523c369587da7293784d49a7502ffcfb0340b1d7885688500ca308161a7f96b62df9d083b71fcc8f2bb8fe6b1689256c0d385f42f5bbe2027a22c1996e110ba97c171d3e5948de92bebfa12649ae8db8acc00a0c4923ad2a4566a7b892f365ad5008fd0aed012ac047f95eec8b2e541cad4e91de38385f2e046619f54496c2382cb6cacd5b98c26f5a4e544041d48115993914bf02fed6cc2d98255a2eb8f91f1e04d5ce16b2306148eef1b63eac20336d5cd32028b1963f7c80869ae34ba13ece0965c51540abc17098a8d7fe3af8caa085a7639a832001457dfb9128a8061142ad0335629ff23ff9cfeb3c337d7a51a6fbf00b9e34c52e1c9195c969bd4e7a0bfd51d5c5bed9c1167e71f0aa83cc32edfbefa9f4d3e0174ca85182eec9f3a09f6a6c0df6377a510d731206fa80a50bb6abe29085058f16212212a60eec8f049fecb92d8c8e0a84bc021352bfecbeddde993839f614c3dac0a3ee37543f9b412b16199dc158e23b544619e312724bb6d7c3153ed9de791d764a366b389af13c58bf8a8d90481a467657cdd2986268250628d0c10e385c58c6191e6fbe05191bcc04f133f2cea72c1c4848930bd7ba8cac54661072113fb278869e07bb8587f91392933374d017bcbe18869ff2c22b28cc10510d9853292803328be4fb0e80495e8bb8d271f5b889636b5fe28e79f1b850f8658246ce9b6a1e7b49fc06db7143e8fe0b4f2b0c5523a5c985e929f70af28d0bdd1a90a808f977f597c7c778c489e98d3bd8910d31ac0f7c6f67e02e6e4e1bdefb994c6098953f34636ba2b6ca20a4721d2b26a886722ff1c9a7e5ff1cf48b4ad1582d3f4e4a1004f3b20d8c5a2b71387a4254ad933ebc52f075ae229646b6f6aed19a5e372cf295081401eb893ff599b3f9acc0c0d3e7d328921deb59612076801e8cd61592107b5c67c79b846595cc6320c395b46362cbfb909fdb236ad2411b4e4883810a074b840464689986c3f8a8091827e17c32755d8fb3687ba3ba49f342c77f5a1f89bec83d811446e1a467139213d640b6a74f7210d4f8e7e1039790e7bf4efa207555a10a6db1dd4b95da313aaa88b88fe76ad21b516cbc645ffe34ab5de1c8aef8cd4e7f8d2b51e8e1456adc7563cda206f9f2c1a00000000000000000000000000000000000000000000000000000000008d36040000000000000000000000000000000000000000000000000000000000983c5974658fd98256d3585175f0be218cda7fba4f73d50eb67f98052f3af700bc2187cfcb64bad6baeceffa46ab6e861e585f0f8fe2999fc171ffced372dbca22da006532a112f2608169a1b29316c8b2467ae6f480161613e7002e684df8179cb92f6003909e7dbc88b4786d5afe771de7547e5df1fba091ddfe88301e100a000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000008175f5626f9bf2906dcc5205aca750c1e96f1833e3455a9126629c2ce1d5f1b17010000000000000000000000fc071a0fd1649241d0b5365eed01aa5ee783b87b0040597307000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000066ca000000000000000000000000000000000000000000000000000000000000eecb000000000000000000000000000000000000000000000000000000000000ffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffffff000000000000000000000000000000000000000000000000"
	//decode
	eigenPodAbi, _ := EigenPodMetaData.GetAbi()
	v := make(map[string]interface{})
	//verifyWithdrawCredential := VerifyWithdrawCredential{}
	decode, err := hex.DecodeString(data)
	if err != nil {
		t.Logf("err:%v", err)
		return
	}
	err = eigenPodAbi.UnpackIntoMap(v, "verifyWithdrawalCredentials", decode[4:])
	if err != nil {
		t.Logf("err:%v", err)
		return
	}
	eigenPodAbi.Methods["verifyWithdrawalCredentials"].Inputs.UnpackIntoMap(v, decode[4:])

	t.Log(v)
	//restakingAbi, _ := Restaking.RestakingMetaData.GetAbi()

}

type VerifyWithdrawCredential struct {
	OracleTimestamp       uint64
	StateRootProof        Restaking.BeaconChainProofsStateRootProof
	ValidatorIndices      []*big.Int
	ValidatorFieldsProofs [][]byte
	ValidatorFields       [][][32]byte
}

func TestRestaking(t *testing.T) {
	data := "3f65cf19000000000000000000000000000000000000000000000000000000006653e86000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000007c0693b21bd86e79438b2872fb7203a95ec456e46b1c3349b783d84637a1769077a000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000608a049ca5e244b93927c28bfe45522b6d96a4e5c6f20567b030b78bed03f5123e5497978b5956205700cb0904c604d7a0d777d77f957c03960c3fcdb052335bf9d781fce6ab374318a8e6b8e812178d230a990a7b4d99a6d42e8056a3dc887986000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000001a1c1e0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000005c0bda5ad5d4261ff9982a9c82102f72454b294d3668aa6873dcd916fbab308b57ed919dddfbaef5c613efb89f58374f0c7f646d21c7ee6b01258fb99bd2df3e8bad5a2e5cafd2b1584313237fb946bbba8e2b8777f04ad2fb7366d26761971e9d87ec495d66ceb2e9afeed45b112f1b6fc7edf39572434f1fb6b620a315431dc3e79520eca463e0bc1dc7f8b6e25b2aeecf95dc94deeb53b8d2f3065599a8ab968ac53dc945ecc755b8a8306f66f28cbd68eb7d0a6641a65fec51964064b773e7d0c2bf61b2b9dd9d556655749169f9f19998cc79fcdceca0483ac44405991e66a76a85f4db1cca79a6eddf8e58d8bd1f42c99c1c99385ccafc3128fdf99b6e1f1e5dd7b3e94a3627828c3c4d34c8ea44ec2e5a9fc943c023e861a3b10fe1b67c87b48ec15f8d8c64574899e31b30e4987931f7e46ec04d908d7db3fcd4935d16a6c952c610c86d53ebed5539f3070ae8770b062f6b97f9317d47877a3d9efcc0b88b2a9c8687497122d1ec6bce6d9a052c193e85c79d16dc9d03596f05dc1188358a7ca2b8279fd0b7f416c389cbc3c2f0d83f39d68e26dc0fff4d296d3a9834d23fc8e359bd64a5f117d3c5b7ee327c88596f098e5e116e89236d18feb4b726ab58d900f5e182e3c50ef74969ea16c7726c549757cc23523c369587da7293784d49a7502ffcfb0340b1d7885688500ca308161a7f96b62df9d083b71fcc8f2bb8fe6b1689256c0d385f42f5bbe2027a22c1996e110ba97c171d3e5948de92bebfa12649ae8db8acc00a0c4923ad2a4566a7b892f365ad5008fd0aed012ac047f95eec8b2e541cad4e91de38385f2e046619f54496c2382cb6cacd5b98c26f5a4e544041d48115993914bf02fed6cc2d98255a2eb8f91f1e04d5ce16b2306148eef1b63eac20336d5cd32028b1963f7c80869ae34ba13ece0965c51540abc17098a8d7fe3af8caa085a7639a832001457dfb9128a8061142ad0335629ff23ff9cfeb3c337d7a51a6fbf00b9e34c52e1c9195c969bd4e7a0bfd51d5c5bed9c1167e71f0aa83cc32edfbefa9f4d3e0174ca85182eec9f3a09f6a6c0df6377a510d731206fa80a50bb6abe29085058f16212212a60eec8f049fecb92d8c8e0a84bc021352bfecbeddde993839f614c3dac0a3ee37543f9b412b16199dc158e23b544619e312724bb6d7c3153ed9de791d764a366b389af13c58bf8a8d90481a467657cdd2986268250628d0c10e385c58c6191e6fbe05191bcc04f133f2cea72c1c4848930bd7ba8cac54661072113fb278869e07bb8587f91392933374d017bcbe18869ff2c22b28cc10510d9853292803328be4fb0e80495e8bb8d271f5b889636b5fe28e79f1b850f8658246ce9b6a1e7b49fc06db7143e8fe0b4f2b0c5523a5c985e929f70af28d0bdd1a90a808f977f597c7c778c489e98d3bd8910d31ac0f7c6f67e02e6e4e1bdefb994c6098953f34636ba2b6ca20a4721d2b26a886722ff1c9a7e5ff1cf48b4ad1582d3f4e4a1004f3b20d8c5a2b71387a4254ad933ebc52f075ae229646b6f6aed19a5e372cf295081401eb893ff599b3f9acc0c0d3e7d328921deb59612076801e8cd61592107b5c67c79b846595cc6320c395b46362cbfb909fdb236ad2411b4e4883810a074b840464689986c3f8a8091827e17c32755d8fb3687ba3ba49f342c77f5a1f89bec83d811446e1a467139213d640b6a74f7210d4f8e7e1039790e7bf4efa207555a10a6db1dd4b95da313aaa88b88fe76ad21b516cbc645ffe34ab5de1c8aef8cd4e7f8d2b51e8e1456adc7563cda206f9f2c1a00000000000000000000000000000000000000000000000000000000008d36040000000000000000000000000000000000000000000000000000000000983c5974658fd98256d3585175f0be218cda7fba4f73d50eb67f98052f3af700bc2187cfcb64bad6baeceffa46ab6e861e585f0f8fe2999fc171ffced372dbca22da006532a112f2608169a1b29316c8b2467ae6f480161613e7002e684df8179cb92f6003909e7dbc88b4786d5afe771de7547e5df1fba091ddfe88301e100a000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000008175f5626f9bf2906dcc5205aca750c1e96f1833e3455a9126629c2ce1d5f1b17010000000000000000000000fc071a0fd1649241d0b5365eed01aa5ee783b87b0040597307000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000066ca000000000000000000000000000000000000000000000000000000000000eecb000000000000000000000000000000000000000000000000000000000000ffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffffff000000000000000000000000000000000000000000000000"
	restakingAbi, _ := Restaking.RestakingMetaData.GetAbi()
	v := make(map[string]interface{})
	eigenPodAbi, _ := EigenPodMetaData.GetAbi()
	//verifyWithdrawCredential := VerifyWithdrawCredential{}
	decode, err := hex.DecodeString(data)
	if err != nil {
		t.Logf("err:%v", err)
		return
	}
	eigenPodAbi.Methods["verifyWithdrawalCredentials"].Inputs.UnpackIntoMap(v, decode[4:])
	//t.Log("v:", v)
	verify := VerifyWithdrawCredential{}
	if value, ok := v["oracleTimestamp"].(uint64); ok {
		verify.OracleTimestamp = value
	}
	if value, ok := v["stateRootProof"].(Restaking.BeaconChainProofsStateRootProof); ok {
		verify.StateRootProof = value
	}
	if value, ok := v["validatorIndices"].([]*big.Int); ok {
		verify.ValidatorIndices = value
	}
	if value, ok := v["validatorFieldsProofs"].([][]byte); ok {
		verify.ValidatorFieldsProofs = value
	}
	if value, ok := v["validatorFields"].([][][32]byte); ok {
		verify.ValidatorFields = value
	}
	//input, err := restakingAbi.Pack("verifyWithdrawalCredentials", common.HexToAddress("0"), v["oracleTimestamp"],
	//	v["stateRootProof"], v["validatorIndices"], v["validatorFieldsProofs"],
	//	v["validatorFields"])
	verifyJson, _ := json.Marshal(&verify)
	t.Logf("verify:%v", string(verifyJson))
	//input, err := restakingAbi.Pack("verifyWithdrawalCredentials", big.NewInt(0), verify.OracleTimestamp,
	//	verify.StateRootProof, verify.ValidatorIndices, verify.ValidatorFieldsProofs,
	//	verify.ValidatorFields)
	input, err := restakingAbi.Pack("verifyWithdrawalCredentials", big.NewInt(0), v["oracleTimestamp"],
		v["stateRootProof"], v["validatorIndices"], v["validatorFieldsProofs"],
		v["validatorFields"])
	if err != nil {
		t.Logf("err:%v", err)
		return
	}
	t.Logf("input:%v", input)
}

func TestParseCredentialProof(t *testing.T) {
	dataHex := "3f65cf190000000000000000000000000000000000000000000000000000000066aa07c000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000007c0775f65b4a7edbaac8638cf3e2c253842ab6e8ceb170a38bd2829e9b3905f422800000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000060493957f2d5e88ff4aaf129cf4acb6654f5c78caeef4a492060534d4f210b4a2f957f2d19f086e5f360f2f753bbea697ff56d75504e297fd3950352dfce80ed3cbcc1402de4947422fcfd68c1cbbffb2e45e1fed28e8085cb95d5d3716ed2c48f000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000001b02420000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000005c0ad555872ff592c98614465db2d3de9989339e778a3083ca509e88144adbb6a1c73408e49ce771fe075870c53d9ec12c36e4d01611e194473c9a92abb4fe901304f3cf9024ffcee72098998b65edc53ed4c9f8248546b2e3cae77f692a16412167b27b70bf859985700d172496cc53721566acc165b2dc2f22c03065ccb1ad1e20fac06b1e69298b5d823d5ebbc0276b9359e55c82b235ea7c5d9c18915843f90cfabdd5f5364cee35c5b34a9bb46a8658459793189c1d2ace471537cb001e0049c61a2746e1e2b67f348584e1efe3b9c69de0a2a2e749103dc56b4ceeadd5a6fcb70a44457fbd8fac4da4087811aa0cb30e3e831ddf92723da986911f46bc13460afa08defcece23be430254b5b2133195c1b8216e9e1086e280b51064428e7f169e16f19a73e0b9b9a5c56adccb7b42c88f276a8d847115292f211dce27c801e924a22d7910254c1f5d8a64a3c4ffe2afa7588a77a58de509d55e6d0f0d821db09de8a3a0af64143eeff6cd4ab39774f4ee9ad3a04d8807d7e7eacc5db6f97935f9d55369f5ba07c5880c790616482b793a26e90a954997ebe60ccef310837cdf6af5f5bbdb6be9ef8aa618e4bf8073960867171e29676f8b284dea6a08a85eb58d900f5e182e3c50ef74969ea16c7726c549757cc23523c369587da7293784d49a7502ffcfb0340b1d7885688500ca308161a7f96b62df9d083b71fcc8f2bbf6547a15e3415494b15ae91fe27e3e92714d148ec745da75c3989bd7aac1d01e52bffe89de70a12c5bd1bf5fa18e29bd97a0a07d3b000d6a993afe51a1a9556a95eec8b2e541cad4e91de38385f2e046619f54496c2382cb6cacd5b98c26f5a4868b18545c5a9d98d38f236447c5c6eb3b269a7be1fba9318abba0994a1fecb8317faa527ddc563edb004e8562a95e53bfe32c6c02eef9715f88caf59b3843c18a8d7fe3af8caa085a7639a832001457dfb9128a8061142ad0335629ff23ff9cfeb3c337d7a51a6fbf00b9e34c52e1c9195c969bd4e7a0bfd51d5c5bed9c1167e71f0aa83cc32edfbefa9f4d3e0174ca85182eec9f3a09f6a6c0df6377a510d731206fa80a50bb6abe29085058f16212212a60eec8f049fecb92d8c8e0a84bc021352bfecbeddde993839f614c3dac0a3ee37543f9b412b16199dc158e23b544619e312724bb6d7c3153ed9de791d764a366b389af13c58bf8a8d90481a467657cdd2986268250628d0c10e385c58c6191e6fbe05191bcc04f133f2cea72c1c4848930bd7ba8cac54661072113fb278869e07bb8587f91392933374d017bcbe18869ff2c22b28cc10510d9853292803328be4fb0e80495e8bb8d271f5b889636b5fe28e79f1b850f8658246ce9b6a1e7b49fc06db7143e8fe0b4f2b0c5523a5c985e929f70af28d0bdd1a90a808f977f597c7c778c489e98d3bd8910d31ac0f7c6f67e02e6e4e1bdefb994c6098953f34636ba2b6ca20a4721d2b26a886722ff1c9a7e5ff1cf48b4ad1582d3f4e4a1004f3b20d8c5a2b71387a4254ad933ebc52f075ae229646b6f6aed19a5e372cf295081401eb893ff599b3f9acc0c0d3e7d328921deb59612076801e8cd61592107b5c67c79b846595cc6320c395b46362cbfb909fdb236ad2411b4e4883810a074b840464689986c3f8a8091827e17c32755d8fb3687ba3ba49f342c77f5a1f89bec83d811446e1a467139213d640b6a74f7210d4f8e7e1039790e7bf4efa207555a10a6db1dd4b95da313aaa88b88fe76ad21b516cbc645ffe34ab5de1c8aef8cd4e7f8d2b51e8e1456adc7563cda206fd81d1b00000000000000000000000000000000000000000000000000000000001b31050000000000000000000000000000000000000000000000000000000000092a3bf3ca5321c88f6642ea96552656b1c8338301c6bc95fe81f8cf88d995c9194d57ae6e660ff0f116fcb5bfb8d406f6bf2442dca54915f07daf46b1f65381b3d2bab771610d7f9253a1a2d910201742ad15428c619978dcf8d3fefda7f42ccd02152d05d0e19d77b91539b69cb91c07883c92f2f50323ac453abcf33eed1f0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000088d7f8dc007d38e6243975d6461156d10146568eda724b398a3204db597b88613010000000000000000000000fc071a0fd1649241d0b5365eed01aa5ee783b87b00405973070000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e50a010000000000000000000000000000000000000000000000000000000000eb0a010000000000000000000000000000000000000000000000000000000000ffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffffff000000000000000000000000000000000000000000000000"
	txData, _ := hex.DecodeString(dataHex)
	eigenPodAbi, _ := EigenPodMetaData.GetAbi()
	m := map[string]interface{}{}
	err := eigenPodAbi.Methods["verifyWithdrawalCredentials"].Inputs.UnpackIntoMap(m, txData[4:])
	if err != nil {
		t.Logf("err:%v", err)
		return
	}
	t.Logf("m:%v", m)
}
func TestGetCurrentCheckPoint(t *testing.T) {
	t.Log("url:", RpcHost)
	cp, err := contract.CurrentCheckpoint(nil)
	if err != nil {
		t.Error(err)
		return
	}
	jsonBytes, _ := json.Marshal(cp)
	t.Log("cp:", string(jsonBytes))
}
