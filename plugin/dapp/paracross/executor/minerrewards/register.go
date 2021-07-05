package minerrewards

import (
	"fmt"

	"github.com/turingchain2020/turingchain/types"
	pt "github.com/turingchain2020/plugin/plugin/dapp/paracross/types"
)

type RewardPolicy interface {
	GetConfigReward(cfg *types.TuringchainConfig, height int64) (int64, int64, int64)
	RewardMiners(cfg *types.TuringchainConfig, coinReward int64, miners []string, height int64) ([]*pt.ParaMinerReward, int64)
}

var MinerRewards = make(map[string]RewardPolicy)

func register(ty string, policy RewardPolicy) {
	if _, ok := MinerRewards[ty]; ok {
		panic(fmt.Sprintf("paracross minerreward ty=%s registered", ty))
	}
	MinerRewards[ty] = policy
}
