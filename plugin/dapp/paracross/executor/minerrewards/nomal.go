package minerrewards

import (
	"github.com/turingchain2020/turingchain/types"
	pt "github.com/turingchain2020/plugin/plugin/dapp/paracross/types"
)

type normal struct{}

func init() {
	register("normal", &normal{})
}

//获取配置的奖励数值
func (n *normal) GetConfigReward(cfg *types.TuringchainConfig, height int64) (int64, int64, int64) {
	coinReward := cfg.MGInt("mver.consensus.paracross.coinReward", height)
	fundReward := cfg.MGInt("mver.consensus.paracross.coinDevFund", height)
	coinBaseReward := cfg.MGInt("mver.consensus.paracross.coinBaseReward", height)

	if coinReward < 0 || fundReward < 0 || coinBaseReward < 0 {
		panic("para config consensus.paracross.coinReward should bigger than 0")
	}

	//decimalMode=false,意味着精简模式，需要乘1e8
	decimalMode := cfg.MIsEnable("mver.consensus.paracross.decimalMode", height)
	if !decimalMode {
		coinReward *= types.Coin
		fundReward *= types.Coin
		coinBaseReward *= types.Coin
	}
	//防止coinBaseReward 设置出错场景， coinBaseReward 一定要比coinReward小
	if coinBaseReward >= coinReward {
		coinBaseReward = coinReward / 10
	}
	return coinReward, fundReward, coinBaseReward
}

//奖励矿工算法
func (n *normal) RewardMiners(cfg *types.TuringchainConfig, coinReward int64, miners []string, height int64) ([]*pt.ParaMinerReward, int64) {
	//找零
	var change int64
	var rewards []*pt.ParaMinerReward
	//分配给矿工的平均奖励
	minerUnit := coinReward / int64(len(miners))
	if minerUnit > 0 {
		for _, m := range miners {
			r := &pt.ParaMinerReward{Addr: m, Amount: minerUnit}
			rewards = append(rewards, r)
		}

		//如果不等分转到发展基金
		change = coinReward % minerUnit
	}
	return rewards, change
}
