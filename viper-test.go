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
	cmds := viper.GetStringMap("testy_manager.cmds")
	testy, _ := cmds["testy"].(map[string]interface{})
	// if !ok {
	// 	return fmt.Errorf("Error parsing testy_manager")
	// }
	cfg := Config{}
	fmt.Println(cfg)
	return &cfg
	return &Config{
		// Directories:    testy["upload_dir"].([]string),
		// SearchInterval: testy["search_interval"].(time.Duration),
		Command: testy["checker_cmd"].(string),
		// Args:           testy["args"].([]string),
	}
}

func main() {
	initConfig("viper-test.json")

	cmds := viper.GetStringMap("testy_manager.cmds")
	testy, _ := cmds["testy"].(map[string]interface{})
	fmt.Println("Directories:", testy["upload_dir"])
	fmt.Println("SearchInterval:", testy["search_interval"])
	fmt.Println("Command:", testy["checker_cmd"])
	fmt.Println("Args:", testy["args"])

	cfg := getTestyManagerConfig()
	fmt.Println("Cfg:", cfg)

}
