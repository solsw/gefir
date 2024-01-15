package chain

import (
	"sync"
)

// well-known ChainId
const (
	ChainIdEthereumMainnet = 1
	ChainIdSepolia         = 11155111
)

var (
	EthereumMainnet func() *Chain
	Sepolia         func() *Chain
)

func init() {
	EthereumMainnet = onceChain(ChainIdEthereumMainnet)
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
