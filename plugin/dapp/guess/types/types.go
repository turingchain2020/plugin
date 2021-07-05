// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"reflect"

	"github.com/turingchain2020/turingchain/types"
)

func init() {
	// init executor type
	types.AllowUserExec = append(types.AllowUserExec, ExecerGuess)
	types.RegFork(GuessX, InitFork)
	types.RegExec(GuessX, InitExecutor)
}

//InitFork ...
func InitFork(cfg *types.TuringchainConfig) {
	cfg.RegisterDappFork(GuessX, "Enable", 0)
}

//InitExecutor ...
func InitExecutor(cfg *types.TuringchainConfig) {
	types.RegistorExecutor(GuessX, NewType(cfg))
}

// GuessType struct
type GuessType struct {
	types.ExecTypeBase
}

// NewType method
func NewType(cfg *types.TuringchainConfig) *GuessType {
	c := &GuessType{}
	c.SetChild(c)
	c.SetConfig(cfg)
	return c
}

// GetPayload method
func (t *GuessType) GetPayload() types.Message {
	return &GuessGameAction{}
}

// GetTypeMap method
func (t *GuessType) GetTypeMap() map[string]int32 {
	return map[string]int32{
		"Start":   GuessGameActionStart,
		"Bet":     GuessGameActionBet,
		"StopBet": GuessGameActionStopBet,
		"Abort":   GuessGameActionAbort,
		"Publish": GuessGameActionPublish,
		"Query":   GuessGameActionQuery,
	}
}

// GetLogMap method
func (t *GuessType) GetLogMap() map[int64]*types.LogInfo {
	return map[int64]*types.LogInfo{
		TyLogGuessGameStart:   {Ty: reflect.TypeOf(ReceiptGuessGame{}), Name: "TyLogGuessGameStart"},
		TyLogGuessGameBet:     {Ty: reflect.TypeOf(ReceiptGuessGame{}), Name: "TyLogGuessGameBet"},
		TyLogGuessGameStopBet: {Ty: reflect.TypeOf(ReceiptGuessGame{}), Name: "TyLogGuessGameStopBet"},
		TyLogGuessGameAbort:   {Ty: reflect.TypeOf(ReceiptGuessGame{}), Name: "TyLogGuessGameAbort"},
		TyLogGuessGamePublish: {Ty: reflect.TypeOf(ReceiptGuessGame{}), Name: "TyLogGuessGamePublish"},
		TyLogGuessGameTimeout: {Ty: reflect.TypeOf(ReceiptGuessGame{}), Name: "TyLogGuessGameTimeout"},
	}
}
