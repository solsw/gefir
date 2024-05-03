package chain

import (
	"testing"
)

func TestChainMiniById(t *testing.T) {
	type args struct {
		chainId uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "0",
			args:    args{chainId: 1234567891011121314},
			wantErr: true,
		},
		{name: "Ethereum Mainnet",
			args: args{chainId: ChainIdEthereumMainnet},
			want: "Ethereum Mainnet",
		},
		{name: "Holesky",
			args: args{chainId: ChainIdHolesky},
			want: "Holesky",
		},
		{name: "Sepolia",
			args: args{chainId: ChainIdSepolia},
			want: "Sepolia",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ChainMiniById(tt.args.chainId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChainMiniById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.Name != tt.want {
				t.Errorf("ChainMiniById().Name = %v, want %v", got.Name, tt.want)
			}
		})
	}
}
