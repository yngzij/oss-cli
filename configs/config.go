package configs

import (
	"github.com/BurntSushi/toml"
)

// Configs is the global config object
type Configs struct {
	Minio MinoConfig `toml:"minio"`
	Psql  PsqlConfig `toml:"psql"`
}

type MinoConfig struct {
	Endpoint        string `toml:"end_point"`
	AccessKeyID     string `toml:"access_key_id"`
	UseSSL          bool   `toml:"use_ssl"`
	SecretAccessKey string `toml:"secret_access_key"`
}

type PsqlConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

// Config is the global config object
var Config *Configs

// LoadConfig loads the config from the given file
func LoadConfig(path string) error {
	if _, err := toml.DecodeFile(path, &Config); err != nil {
		return err
	}

	return nil
}
