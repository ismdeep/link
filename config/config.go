package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type config struct {
	Server struct {
		Bind string
		Site string
	}
	DB struct {
		Dialect string
		DSN     string
	}
}

// WorkDir working directory
var WorkDir string

// Config instance
var Config *config

func init() {
	WorkDir = os.Getenv("LINK_WORK_DIR")
	if WorkDir == "" {
		WorkDir, _ = os.Getwd()
	}

	content, err := ioutil.ReadFile(fmt.Sprintf("%v/config.toml", WorkDir))
	if err != nil {
		panic(err)
	}

	Config = &config{}
	if err := toml.Unmarshal(content, Config); err != nil {
		panic(err)
	}
}
