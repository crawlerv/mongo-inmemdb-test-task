package resources

import (
	"flag"
	"fmt"
	"github.com/crawlerv/mongo-inmemdb-test-task/config"
	"github.com/spf13/cobra"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
	once    sync.Once
	cfg     *config.Config
	conFile string
)

func GetConfig() *config.Config {
	once.Do(func() {
		if conFile == "" {
			panic("config file not set")
		}
		viper.SetConfigFile(strings.TrimSpace(conFile))
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Sprintf("config not loaded, %v", err))
		}
		err := viper.Unmarshal(&cfg)
		if err != nil {
			panic(fmt.Sprintf("unable to decode into config struct, %v", err))
		}
	})
	return cfg
}

func RegisterConfigFlag(c *cobra.Command) {
	c.PersistentFlags().StringVarP(&conFile, "config", "c", "", `Path to config file. Supports .json, .yaml, .yml, .toml.`)
	c.PersistentFlags().AddGoFlagSet(flag.CommandLine)
}
