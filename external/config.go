package external

import (
	"context"
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	ServerHost     string
	DatabaseConfig DatabaseConfig
	Environment    string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

var (
	runOnce sync.Once
	config  Config
)

func GetConfig() Config {
	runOnce.Do(func() {
		cfg, err := initConfig()
		if err != nil {
			fmt.Println(context.Background(), err, "could not load usecase configuration")
		}

		fmt.Println(cfg.GetString("environment"))
		config = Config{
			ServerHost: cfg.GetString("server.host"),
			DatabaseConfig: DatabaseConfig{
				Host:     cfg.GetString("database.host"),
				Port:     cfg.GetString("database.port"),
				User:     cfg.GetString("database.user"),
				Password: cfg.GetString("database.password"),
				DbName:   cfg.GetString("database.dbname"),
			},
			Environment: cfg.GetString("environment"),
		}
	})

	return config
}

func initConfig() (viper.Viper, error) {
	cfg := viper.New()
	cfg.AutomaticEnv()
	var err error
	initDefaults(cfg)
	// workaround because viper does not resolve envs when unmarshalling
	for _, key := range cfg.AllKeys() {
		val := cfg.Get(key)
		cfg.Set(key, val)
	}
	return *cfg, err
}

func initDefaults(config *viper.Viper) {
	config.SetDefault("server.host", "0.0.0.0:8000")
	config.SetDefault("database.host", "postgres")
	config.SetDefault("database.port", "5432")
	config.SetDefault("database.user", "root")
	config.SetDefault("database.password", "root")
	config.SetDefault("database.dbname", "root")
	config.SetDefault("environment", "production")
}
