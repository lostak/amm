package keeper

import (
	"github.com/lostak/amm/x/ocean/types"
)

var _ types.QueryServer = Keeper{}
