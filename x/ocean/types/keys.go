package types

import (
	"fmt"
	"strings"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "ocean"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ocean"

	// Version defines the current version the IBC module supports
	Version = "ocean-1"

	// PortID is the default port id that module binds to
	PortID = "ocean"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("ocean-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func PoolIndex(tokenA sdk.Coin,tokenB sdk.Coin, ) (string, error) {

	denomA := strings.ToUpper(tokenA.Denom)
	denomB := strings.ToUpper(tokenB.Denom)

	if denomA == denomB {
		return "", errors.New("coins have same denoms")
	}

	if denomA > denomB {
		return fmt.Sprintf("%s-%s", denomA, denomB, ), nil
	}

	return fmt.Sprintf("%s-%s", denomB, denomA, ), nil
}
