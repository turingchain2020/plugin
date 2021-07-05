/*Package commands implement dapp client commands*/
package commands

import (
	jsonrpc "github.com/turingchain2020/turingchain/rpc/jsonclient"
	rpctypes "github.com/turingchain2020/turingchain/rpc/types"
	"github.com/turingchain2020/turingchain/types"
	vty "github.com/turingchain2020/plugin/plugin/dapp/vote/types"
	"github.com/spf13/cobra"
)

/*
 * 实现合约对应客户端
 */

// Cmd vote client command
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote",
		Short: "vote command",
		Args:  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(
		//create tx
		createGroupCMD(),
		updateGroupCMD(),
		createVoteCMD(),
		commitVoteCMD(),
		closeVoteCMD(),
		updateMemberCMD(),
		//query rpc
		groupInfoCMD(),
		voteInfoCMD(),
		memberInfoCMD(),
		listGroupCMD(),
		listVoteCMD(),
		listMemberCMD(),
	)
	return cmd
}

func markRequired(cmd *cobra.Command, params ...string) {
	for _, param := range params {
		_ = cmd.MarkFlagRequired(param)
	}
}

func sendCreateTxRPC(cmd *cobra.Command, actionName string, req types.Message) {
	title, _ := cmd.Flags().GetString("title")
	cfg := types.GetCliSysParam(title)
	rpcAddr, _ := cmd.Flags().GetString("rpc_laddr")
	payLoad := types.MustPBToJSON(req)
	pm := &rpctypes.CreateTxIn{
		Execer:     cfg.ExecName(vty.VoteX),
		ActionName: actionName,
		Payload:    payLoad,
	}

	var res string
	ctx := jsonrpc.NewRPCCtx(rpcAddr, "Turingchain.CreateTransaction", pm, &res)
	ctx.RunWithoutMarshal()
}

func sendQueryRPC(cmd *cobra.Command, funcName string, req, reply types.Message) {
	title, _ := cmd.Flags().GetString("title")
	cfg := types.GetCliSysParam(title)
	rpcAddr, _ := cmd.Flags().GetString("rpc_laddr")
	payLoad := types.MustPBToJSON(req)
	query := &rpctypes.Query4Jrpc{
		Execer:   cfg.ExecName(vty.VoteX),
		FuncName: funcName,
		Payload:  payLoad,
	}

	ctx := jsonrpc.NewRPCCtx(rpcAddr, "Turingchain.Query", query, reply)
	ctx.Run()
}
