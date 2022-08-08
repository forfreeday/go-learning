package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
)
import config "golearning/src/p24toml"

var (
	cfg  *config.TomlConfig
	once sync.Once
)

func Config() *config.TomlConfig {
	once.Do(func() {
		filePath, err := filepath.Abs("/Users/liukai/workspaces/go/my-project/go-learning/src/p24toml/toml/config.toml")
		if err != nil {
			panic(err)
		}
		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}

func main() {
	tomlConfig := Config()
	fmt.Println(tomlConfig.Title)
	// 配置中DB的IP
	fmt.Println(tomlConfig.DB.Server)
	// 配置中Owner的名字
	fmt.Println(tomlConfig.Owner.Name)

}
