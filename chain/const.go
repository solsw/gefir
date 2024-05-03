package chain

import (
	"sync"
)

// well-known ChainId
const (
	ChainIdEthereumMainnet = 1
	ChainIdHolesky         = 17000
	ChainIdSepolia         = 11155111
)

var (
	EthereumMainnet func() *Chain
	Holesky         func() *Chain
	Sepolia         func() *Chain
)

func init() {
	EthereumMainnet = onceChain(ChainIdEthereumMainnet)
	Holesky = onceChain(ChainIdHolesky)
	Sepolia = onceChain(ChainIdSepolia)
}

func onceChain(chainId uint64) func() *Chain {
	return sync.OnceValue(func() *Chain {
		c, err := ChainById(chainId)
		if err != nil {
			panic(err)
		}
		return c
	})
}
