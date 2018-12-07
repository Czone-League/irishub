package protocol

import (
	sdk "github.com/irisnet/irishub/types"
	"github.com/irisnet/irishub/types/common"
)

type Protocol interface {
	GetDefinition() common.ProtocolDefinition
	GetRouter() Router
	GetQueryRouter() QueryRouter
	GetAnteHandler() sdk.AnteHandler                   // ante handler for fee and auth
	GetFeeRefundHandler() sdk.FeeRefundHandler         // fee handler for fee refund
	GetFeePreprocessHandler() sdk.FeePreprocessHandler // fee handler for fee preprocessor

	// may be nil
	GetInitChainer() sdk.InitChainer1  // initialize state with validators and state blob
	GetBeginBlocker() sdk.BeginBlocker // logic to run before any txs
	GetEndBlocker() sdk.EndBlocker     // logic to run after all txs, and to determine valset changes
	Load()
	Init()
}

type ProtocolBase struct {
	Definition common.ProtocolDefinition
	//	engine 		*ProtocolEngine
}

func (pb ProtocolBase) GetDefinition() common.ProtocolDefinition {
	return pb.Definition
}