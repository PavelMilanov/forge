package config

import (
	"fmt"

	"github.com/spf13/viper"
)

/*
Env описывает конфигурацию приложения.
*/
type Env struct {
	Vault    vault
	Registry registry
	SSH      ssh
}

/*
vault описывает конфигурацию Hashicorp Vault.
*/
type vault struct {
	Url   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
}

/*
registry описывает конфигурацию Docker Registry.
*/
type registry struct {
	Url      string `mapstructure:"url"`
	Login    string `mapstructure:"login"`
	Password string `mapstructure:"password"`
}

/*
ssh описывает конфигурацию SSH.
*/
type ssh struct {
	PrivateKey string `mapstructure:"private_key"`
	KnownHosts string `mapstructure:"known_hosts"`
}

/*
NewEnv создает экземпляр Env из файла конфигурации.

Returns

	*Env - экземпляр Env
	error - ошибка при создании экземпляра Env
*/
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
	if env.Vault.Url == "" || env.Vault.Token == "" {
		return nil, fmt.Errorf("invalid vault configuration")
	}
	return &env, nil
}
