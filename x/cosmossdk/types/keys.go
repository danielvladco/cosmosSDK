package types

const (
	// ModuleName defines the module name
	ModuleName = "cosmossdk"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cosmossdk"
)

var (
	ParamsKey = []byte("p_cosmossdk")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
