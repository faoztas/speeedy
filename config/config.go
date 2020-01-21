package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type env struct {
	Telegram struct {
		Token string
	}
	Logging logging
	DBFile  string
}

type logging struct {
	Level  string
	Format string
}

var (
	Env env
)

func Init(path string) {
	var envPath string
	if path == "" {
		if envPath = os.Getenv("SPEEDTEST_CONFIG_FILE_PATH"); envPath == "" {
			panic("config file path must be gived with flag or SPEEDTEST_CONFIG_FILE_PATH env var")
		}
	}

	if _, err := toml.DecodeFile(path, &Env); err != nil {
		panic(fmt.Errorf("config file read error : %e", err))
	}

	if Env.DBFile == "" {
		Env.DBFile = "ledger.db"
	}

}
