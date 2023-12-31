package cmd

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var (
	validate           = validator.New()
	ErrEmptyConfigPath = errors.New("empty configuration file path")
)

type (
	Config struct {
		TimeoutHeight int64          `mapstructure:"timeout_height" validate:"required"`
		EthConfig     EthChainConfig `mapstructure:"eth_chain_config" validate:"dive,required"`
		GasConfig     Gas            `mapstructure:"gas_config" validate:"required"`
		Account       Account        `mapstructure:"account" validate:"required,gt=0,dive,required"`
		Keyring       Keyring        `mapstructure:"keyring" validate:"required,gt=0,dive,required"`
		RPC           RPC            `mapstructure:"rpc" validate:"required,gt=0,dive,required"`
	}

	EthChainConfig struct {
		ContractAddress string `mapstructure:"contract_address" validate:"required"`
		RPC             string `mapstructure:"rpc" validate:"required,required"`
	}

	Gas struct {
		GasPrice      string  `mapstructure:"gas_price" validate:"required"`
		GasAdjustment float64 `mapstructure:"gas_adjustment" validate:"required"`
	}

	Account struct {
		ChainID string `mapstructure:"chain_id" validate:"required"`
		Address string `mapstructure:"address" validate:"required"`
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

func (c Config) Validate() error {
	return validate.Struct(c)
}

func ParseConfig(configPath string) (Config, error) {
	var cfg Config

	if configPath == "" {
		return cfg, ErrEmptyConfigPath
	}

	viper.AutomaticEnv()
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("failed to read config: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to decode config: %w", err)
	}

	return cfg, cfg.Validate()
}
