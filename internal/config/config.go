package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"PORT"`
	Host        string `mapstructure:"HOST"`
	TLSCertFile string `mapstructure:"TLS_CERT_FILE"`
	TLSKeyFile  string `mapstructure:"TLS_KEY_FILE"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
