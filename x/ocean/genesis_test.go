package ocean_test

import (
	"testing"

	keepertest "github.com/lostak/amm/testutil/keeper"
	"github.com/lostak/amm/testutil/nullify"
	"github.com/lostak/amm/x/ocean"
	"github.com/lostak/amm/x/ocean/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		PoolList: []types.Pool{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OceanKeeper(t)
	ocean.InitGenesis(ctx, *k, genesisState)
	got := ocean.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.PoolList, got.PoolList)
	// this line is used by starport scaffolding # genesis/test/assert
}
