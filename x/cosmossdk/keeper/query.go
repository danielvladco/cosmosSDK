package keeper

import (
	"github.com/danielvladco/cosmosSDK/x/cosmossdk/types"
)

var _ types.QueryServer = Keeper{}
