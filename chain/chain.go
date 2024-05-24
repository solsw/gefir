package chain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
	"github.com/solsw/httphelper"
	"github.com/solsw/jsonhelper"
	"github.com/solsw/oshelper"
)

const (
	chainsUrl      = "https://chainid.network/chains.json"
	chainsJson     = "chains.json"
	chainsMiniUrl  = "https://chainid.network/chains_mini.json"
	chainsMiniJson = "chains_mini.json"
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

func getUrlAndName[C ChainMini | Chain]() (url string, name string) {
	var c0 C
	switch any(c0).(type) {
	case ChainMini:
		return chainsMiniUrl, chainsMiniJson
	case Chain:
		return chainsUrl, chainsJson
	default:
		return "", ""
	}
}

// getChains returns slice of [ChainMini] or [Chain].
func getChains[C ChainMini | Chain](dir string) ([]C, error) {
	url, name := getUrlAndName[C]()
	fullname := filepath.Join(dir, name)
	exists, err := oshelper.FileExists(fullname)
	if err != nil {
		return nil, err
	}
	if !exists {
		if err := getJson(url, fullname); err != nil {
			return nil, err
		}
	}
	bb, err := os.ReadFile(fullname)
	if err != nil {
		return nil, err
	}
	var cc []C
	if err := json.Unmarshal(bb, &cc); err != nil {
		return nil, err
	}
	return cc, nil
}

// ChainsDir returns slice of [Chain].
//
// Chains data are loaded from 'chains.json' file in 'dir' directory
// (if it does not exist, it is downloaded from 'https://chainid.network/chains.json').
// To renew 'chains.json' file, remove it.
func ChainsDir(dir string) ([]Chain, error) {
	cc, err := getChains[Chain](dir)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return cc, nil
}

// Chains returns slice of [Chain].
//
// Chains data are loaded from 'chains.json' file in current directory
// (if it does not exist, it is downloaded from 'https://chainid.network/chains.json').
// To renew 'chains.json' file, remove it.
func Chains() ([]Chain, error) {
	cc, err := getChains[Chain](".")
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return cc, nil
}

// ChainsMiniDir returns slice of [ChainMini].
//
// ChainsMini data are loaded from 'chains_mini.json' file in 'dir' directory
// (if it does not exist, it is downloaded from 'https://chainid.network/chains_mini.json').
// To renew 'chains_mini.json' file, remove it.
func ChainsMiniDir(dir string) ([]ChainMini, error) {
	cc, err := getChains[ChainMini](dir)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return cc, nil
}

// ChainsMini returns slice of [ChainMini].
//
// ChainsMini data are loaded from 'chains_mini.json' file in current directory
// (if it does not exist, it is downloaded from 'https://chainid.network/chains_mini.json').
// To renew 'chains_mini.json' file, remove it.
func ChainsMini() ([]ChainMini, error) {
	cc, err := getChains[ChainMini](".")
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return cc, nil
}

// ChainById returns [Chain] by 'chainId'.
//
// Chains data ('chains.json' file) are stored in current directory.
func ChainById(chainId uint64) (*Chain, error) {
	cc, err := Chains()
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	i := slices.IndexFunc(cc, func(c Chain) bool { return c.ChainId == chainId })
	if i < 0 {
		return nil, errorhelper.CallerError(fmt.Errorf("chainId '%d' not found", chainId))
	}
	// to not escape the whole slice to the heap
	c := cc[i]
	return &c, nil
}

// ChainMiniById returns [ChainMini] by 'chainId'.
//
// ChainsMini data ('chains_mini.json' file) are stored in current directory.
func ChainMiniById(chainId uint64) (*ChainMini, error) {
	cc, err := ChainsMini()
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	i := slices.IndexFunc(cc, func(c ChainMini) bool { return c.ChainId == chainId })
	if i < 0 {
		return nil, errorhelper.CallerError(fmt.Errorf("chainId '%d' not found", chainId))
	}
	// to not escape the whole slice to the heap
	c := cc[i]
	return &c, nil
}
