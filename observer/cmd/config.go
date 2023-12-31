package cmd

type (
	Config struct {
		ConfigDir string         `mapstructure:"config_dir"`
		EthConfig EthChainConfig `mapstructure:"eth_chain_config" validate:"dive,required"`
		Account   Account        `mapstructure:"account" validate:"required,gt=0,dive,required"`
		Keyring   Keyring        `mapstructure:"keyring" validate:"required,gt=0,dive,required"`
		RPC       RPC            `mapstructure:"rpc" validate:"required,gt=0,dive,required"`
	}

	EthChainConfig struct {
		ContractAddress string `mapstructure:"contract_address" validate:"required"`
		RPC             string `mapstructure:"rpc" validate:"required,required"`
	}

	Account struct {
		ChainID   string `mapstructure:"chain_id" validate:"required"`
		Address   string `mapstructure:"address" validate:"required"`
		Validator string `mapstructure:"validator" validate:"required"`
	}

	Keyring struct {
		Backend string `mapstructure:"backend" validate:"required"`
		Dir     string `mapstructure:"dir" validate:"required"`
	}

	RPC struct {
		TMRPCEndpoint string `mapstructure:"tmrpc_endpoint" validate:"required"`
		GRPCEndpoint  string `mapstructure:"grpc_endpoint" validate:"required"`
		RPCTimeout    string `mapstructure:"rpc_timeout" validate:"required"`
	}
)
