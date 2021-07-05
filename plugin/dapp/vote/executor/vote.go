package executor

import (
	log "github.com/turingchain2020/turingchain/common/log/log15"
	drivers "github.com/turingchain2020/turingchain/system/dapp"
	"github.com/turingchain2020/turingchain/types"
	votetypes "github.com/turingchain2020/plugin/plugin/dapp/vote/types"
)

/*
 * 执行器相关定义
 * 重载基类相关接口
 */

var (
	//日志
	elog = log.New("module", "vote.executor")
)

var driverName = votetypes.VoteX

// Init register dapp
func Init(name string, cfg *types.TuringchainConfig, sub []byte) {
	drivers.Register(cfg, GetName(), newVote, cfg.GetDappFork(driverName, "Enable"))
	InitExecType()
}

// InitExecType Init Exec Type
func InitExecType() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&vote{}))
}

type vote struct {
	drivers.DriverBase
}

func newVote() drivers.Driver {
	t := &vote{}
	t.SetChild(t)
	t.SetExecutorType(types.LoadExecutorType(driverName))
	return t
}

// GetName get driver name
func GetName() string {
	return newVote().GetName()
}

func (v *vote) GetDriverName() string {
	return driverName
}
