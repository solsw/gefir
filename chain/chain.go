package chain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/solsw/generichelper"
	"github.com/solsw/httphelper"
	"github.com/solsw/jsonhelper"
	"github.com/solsw/oshelper"
)

const (
	chainsUrl      = "https://chainid.network/chains.json"
	chainsName     = "chains.json"
	chainsMiniUrl  = "https://chainid.network/chains_mini.json"
	chainsMiniName = "chains_mini.json"
)

func getJson(url, name string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		herr, err := httphelper.NewError[generichelper.NoType](res)
		if err != nil {
			return err
		}
		return herr
	}
	i, err := jsonhelper.IndentStr(string(body), "", "\t")
	if err != nil {
		return err
	}
	if err := os.WriteFile(name, []byte(i), os.ModePerm); err != nil {
		return err
	}
	return nil
}

func getChains[C ChainMini | Chain](url, name string) ([]C, error) {
	exists, err := oshelper.FileExists(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		if err := getJson(url, name); err != nil {
			return nil, err
		}
	}
	bb, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	var cc []C
	if err := json.Unmarshal(bb, &cc); err != nil {
		return nil, err
	}
	return cc, nil
}

// Chains returns slice of [Chain].
func Chains() ([]Chain, error) {
	return getChains[Chain](chainsUrl, chainsName)
}

// ChainsMini returns slice of [ChainMini].
func ChainsMini() ([]ChainMini, error) {
	return getChains[ChainMini](chainsMiniUrl, chainsMiniName)
}

// ChainById returns [Chain] by 'chainId'.
func ChainById(chainId uint64) (*Chain, error) {
	cc, err := Chains()
	if err != nil {
		return nil, err
	}
	for _, c := range cc {
		if c.ChainId == chainId {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("chainId '%d' not found", chainId)
}

// ChainMiniById returns [ChainMini] by 'chainId'.
func ChainMiniById(chainId uint64) (*ChainMini, error) {
	cc, err := ChainsMini()
	if err != nil {
		return nil, err
	}
	for _, c := range cc {
		if c.ChainId == chainId {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("chainId '%d' not found", chainId)
}
