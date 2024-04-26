package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	config "viper-tool/cfg"
)

func main() {
	viper.SetConfigFile("cfg/cfg.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	viper.SetEnvPrefix("SCAN")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic("resolve to config structure err")
	}
	_, _path, _, _ := runtime.Caller(0)
	fmt.Println("project root dir", filepath.Join(filepath.Dir(_path), ".."))
}
