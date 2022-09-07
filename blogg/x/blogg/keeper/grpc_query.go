package keeper

import (
	"blogg/x/blogg/types"
)

var _ types.QueryServer = Keeper{}
