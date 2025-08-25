package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Env описывает конфигурацию приложения.
type Env struct {
	Vault    vault
	Registry registry
}

// server описывает конфигурацию сервера.
type vault struct {
	Url   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
	Path  string `mapstructure:"path"`
}

// registry описывает конфигурацию хранилища.
type registry struct {
	Url      string `mapstructure:"url"`
	Login    string `mapstructure:"login"`
	Password string `mapstructure:"password"`
}

func NewEnv(path, file string) (*Env, error) {
	env := Env{}
	viper.SetConfigName(file) // имя файла без расширения
	viper.SetConfigType("yml")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		return nil, err
	}
	if env.Vault.Url == "" || env.Vault.Token == "" || env.Vault.Path == "" {
		return nil, fmt.Errorf("invalid vault configuration")
	}
	return &env, nil
}
