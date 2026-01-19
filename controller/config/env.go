package config

import (
	"fmt"

	"github.com/spf13/viper"
)

/*
Env описывает конфигурацию приложения.
*/
type Env struct {
	Vault     vault
	Registry  registry
	Portainer portainer
}

/*
vault описывает конфигурацию Hashicorp Vault.
*/
type vault struct {
	Url string `mapstructure:"url"`
	// Token string `mapstructure:"token"`
	Role   string `mapstructure:"role_id"`
	Secret string `mapstructure:"secret_id"`
}

/*
registry описывает конфигурацию Docker Registry.
*/
type registry struct {
	Url      string `mapstructure:"url"`
	Login    string `mapstructure:"login"`
	Password string `mapstructure:"password"`
}

type portainer struct {
	Url   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
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
	if env.Vault.Url == "" {
		return nil, fmt.Errorf("No parse vault.url")
	}
	if env.Vault.Role == "" || env.Vault.Secret == "" {
		return nil, fmt.Errorf("No parse vault.role_id or vault.secret_id")
	}
	return &env, nil
}
