package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	SubModuleName = "furyint-client"
)

// IBC furyint client sentinel errors
var (
	ErrInvalidChainID          = sdkerrors.Register(SubModuleName, 2, "invalid chain-id")
	ErrInvalidTrustingPeriod   = sdkerrors.Register(SubModuleName, 3, "invalid trusting period")
	ErrInvalidHeaderHeight     = sdkerrors.Register(SubModuleName, 4, "invalid header height")
	ErrInvalidHeader           = sdkerrors.Register(SubModuleName, 5, "invalid header")
	ErrInvalidMaxClockDrift    = sdkerrors.Register(SubModuleName, 6, "invalid max clock drift")
	ErrProcessedTimeNotFound   = sdkerrors.Register(SubModuleName, 7, "processed time not found")
	ErrProcessedHeightNotFound = sdkerrors.Register(SubModuleName, 8, "processed height not found")
	ErrDelayPeriodNotPassed    = sdkerrors.Register(SubModuleName, 9, "packet-specified delay period has not been reached")
	ErrTrustingPeriodExpired   = sdkerrors.Register(SubModuleName, 10, "time since latest trusted state has passed the trusting period")
	ErrInvalidProofSpecs       = sdkerrors.Register(SubModuleName, 12, "invalid proof specs")
	ErrInvalidValidatorSet     = sdkerrors.Register(SubModuleName, 13, "invalid validator set")
)
