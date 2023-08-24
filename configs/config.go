package configs

import (
	"github.com/urfave/cli/v2"
)

var Cfg *Config = &Config{
	//todo
}

type Config struct {
	//todo config filed
}

func Load(ctx *cli.Context) error {
	//todo load args
	return nil
}
