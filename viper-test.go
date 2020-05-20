package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config defines the testy manager configuration
type Config struct {
	Directories    []string
	SearchInterval time.Duration
	Command        string
	Args           []string
}

func initConfig(conf string) error {
	viper.SetConfigFile(conf)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Params file %s not found", conf)
	}
	return nil
}

func getTestyManagerConfigs() []Config {
	var configs []Config
	for key := range viper.GetStringMap("testy_manager.cmds") {
		cfg := Config{
			Directories:    viper.GetStringSlice("testy_manager.cmds." + key + ".upload_dir"),
			SearchInterval: viper.GetDuration("testy_manager.cmds." + key + ".search_interval"),
			Command:        viper.GetString("testy_manager.cmds." + key + ".checker_cmd"),
			Args:           viper.GetStringSlice("testy_manager.cmds." + key + ".args"),
		}
		configs = append(configs, cfg)
	}
	return configs
}

func main() {
	initConfig("viper-test.json")

	cfgs := getTestyManagerConfigs()
	fmt.Println("Cfgs:", cfgs[0])

}
