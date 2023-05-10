package furyint

import (
	"github.com/cosmos/ibc-go/v3/modules/light-clients/01-furyint/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
