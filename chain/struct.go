package chain

// https://github.com/ethereum-lists/chains

// https://docs.infura.io/infura-expansion-apis/gas-api/supported-networks
// https://chainlist.org/
// https://chainid.network/
// https://chainid.network/chains.json
// https://chainid.network/chains_mini.json

type Currency struct {
	Decimals uint64 `json:"decimals"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
}

type ChainMini struct {
	ChainId        uint64   `json:"chainId"`
	Faucets        []string `json:"faucets"`
	InfoURL        string   `json:"infoURL"`
	Name           string   `json:"name"`
	NativeCurrency Currency `json:"nativeCurrency"`
	NetworkId      uint64   `json:"networkId"`
	Rpc            []string `json:"rpc"`
	ShortName      string   `json:"shortName"`
}

// Ethereum Name Service
type Ens struct {
	Registry string `json:"registry"`
}

type Explorer struct {
	Icon     string `json:"icon,omitempty"`
	Name     string `json:"name"`
	Standard string `json:"standard"`
	Url      string `json:"url"`
}

type Feature struct {
	Name string `json:"name"`
}

type Bridge struct {
	Url string `json:"url"`
}

type Parent struct {
	Bridges []Bridge `json:"bridges,omitempty"`
	Chain   string   `json:"chain"`
	Type    string   `json:"type"`
}

type Chain struct {
	ChainMini
	Chain     string     `json:"chain"`
	Ens       Ens        `json:"ens,omitempty"`
	Explorers []Explorer `json:"explorers,omitempty"`
	Features  []Feature  `json:"features,omitempty"`
	Icon      string     `json:"icon,omitempty"`
	Parent    Parent     `json:"parent,omitempty"`
	Slip44    uint64     `json:"slip44,omitempty"`
	Status    string     `json:"status,omitempty"`
}
