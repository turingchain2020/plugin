package para

import (
	"github.com/turingchain2020/turingchain/queue"
	drivers "github.com/turingchain2020/turingchain/system/mempool"
	"github.com/turingchain2020/turingchain/types"
)

//--------------------------------------------------------------------------------
// Module Mempool

func init() {
	drivers.Reg("para", New)
}

//New 创建price cache 结构的 mempool
func New(cfg *types.Mempool, sub []byte) queue.Module {
	return NewMempool(cfg)
}
