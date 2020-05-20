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

func getTestyManagerConfig() *Config {
	return &Config{
		Directories:    viper.GetStringSlice("testy_manager.upload_dir"),
		SearchInterval: viper.GetDuration("testy_manager.search_interval"),
		Command:        viper.GetString("testy_manager.checker_cmd"),
		Args:           viper.GetStringSlice("testy_manager.args"),
	}
}

func main() {
	initConfig("viper-test.json")
	// cfg := getTestyManagerConfig()
	cmds := viper.Get("testy_manager")
	fmt.Println("Cmds:", cmds)
}
