package rand

import (
	"github.com/irisnet/irishub/app/v3/rand/internal/keeper"
	"github.com/irisnet/irishub/app/v3/rand/internal/types"
)

// exported types
type (
	MsgRequestRand              = types.MsgRequestRand
	Rand                        = types.Rand
	Request                     = types.Request
	Requests                    = types.Requests
	GenesisState                = types.GenesisState
	QueryRandParams             = types.QueryRandParams
	QueryRandRequestQueueParams = types.QueryRandRequestQueueParams
	Keeper                      = keeper.Keeper
)

// exported consts
const (
	ServiceName           = types.ModuleName
	DefaultCodespace      = types.DefaultCodespace
	DefaultBlockInterval  = types.DefaultBlockInterval
	RandPrec              = types.RandPrec
	QueryRand             = types.QueryRand
	QueryRandRequestQueue = types.QueryRandRequestQueue
)

// exported variables and functions
var (
	TagReqID            = types.TagReqID
	TagRequestContextID = types.TagRequestContextID
	TagRandHeight       = types.TagRandHeight
	TagRand             = types.TagRand

	RegisterCodec     = types.RegisterCodec
	NewMsgRequestRand = types.NewMsgRequestRand
	NewRand           = types.NewRand
	NewRequest        = types.NewRequest
	MakePRNG          = types.MakePRNG
	GenerateRequestID = types.GenerateRequestID
	CheckReqID        = types.CheckReqID
	NewKeeper         = keeper.NewKeeper
	NewQuerier        = keeper.NewQuerier
)