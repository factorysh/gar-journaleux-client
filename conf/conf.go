package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type ProjectInfo struct {
	Name string
}

type server struct {
	Address string
}

type ClientConfig struct {
	Project ProjectInfo
	Servers map[string]server
}

func ReadClientConf(conf *ClientConfig, path string) error {

	if _, err := os.Stat(path); err == nil {
		// do stuff only if file exists
		_, err := toml.DecodeFile(path, conf)
		return err
	}

	fmt.Println(conf)

	return nil
}
