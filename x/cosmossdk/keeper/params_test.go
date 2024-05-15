package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/danielvladco/cosmosSDK/testutil/keeper"
	"github.com/danielvladco/cosmosSDK/x/cosmossdk/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CosmossdkKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
