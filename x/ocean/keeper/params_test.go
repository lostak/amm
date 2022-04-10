package keeper_test

import (
	"testing"

	testkeeper "github.com/lostak/amm/testutil/keeper"
	"github.com/lostak/amm/x/ocean/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OceanKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
