package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"github.com/wavesplatform/GatewaysInfrastructure/Adapters/Ergo/logger"
	"github.com/wavesplatform/GatewaysInfrastructure/Adapters/Ergo/models"
)

const (
	LogLevelEnvKey  = "LOG_LEVEL"
	LogLevelDefault = logger.INFO
)

var Cfg *Config

type Config struct {
	Node Node   `mapstructure:"NODE"`
	Port string `mapstructure:"PORT"`
}

type Node struct {
	Host        string             `mapstructure:"HOST"`
	ExplorerUrl string             `mapstructure:"EXPLORER_URL"`
	ChainID     models.NetworkType `mapstructure:"CHAINID"`
	ApiKey      string             `mapstructure: "APIKEY"`
}

func (c *Config) String() string {
	return fmt.Sprintf("NODE_HOST: %s, NODE_CHAINID: %v, PORT: %s",
		c.Node.Host, c.Node.ChainID, c.Port,
	)
}

// Load set configuration parameters.
// At first read config from file
// After that read environment variables
func Load(defaultConfigPath string) error {
	cfg := new(Config)

	// read config from file - it will be default values
	viper.SetConfigFile(defaultConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}

	// read parameters from environment variables -> they override default values from file
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}
	Cfg = cfg
	return validateConfig()
}

func validateConfig() error {
	if len(Cfg.Node.Host) == 0 {
		return errors.New("NODE_HOST parameter is empty")
	}
	if len(Cfg.Port) == 0 {
		return errors.New("PORT parameter is empty")
	}
	return nil
}
