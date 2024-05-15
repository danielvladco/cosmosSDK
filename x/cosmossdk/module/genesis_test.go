package cosmossdk_test

import (
	"testing"

	keepertest "github.com/danielvladco/cosmosSDK/testutil/keeper"
	"github.com/danielvladco/cosmosSDK/testutil/nullify"
	cosmossdk "github.com/danielvladco/cosmosSDK/x/cosmossdk/module"
	"github.com/danielvladco/cosmosSDK/x/cosmossdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmossdkKeeper(t)
	cosmossdk.InitGenesis(ctx, k, genesisState)
	got := cosmossdk.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
