package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

var AppConfig *Env

/*
Env описывает конфигурацию приложения.
*/
type Env struct {
	Vault    vault
	Registry registry
	Agent    agent
	// Portainer portainer
}

/*
vault описывает конфигурацию Hashicorp Vault.
*/
type vault struct {
	Url    string `mapstructure:"url"`
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
	Url string `mapstructure:"url"`
	Key string `mapstructure:"key"`
}

type agent struct {
	Type        string    `mapstructure:"type"`
	Credentials portainer `mapstructure:"credentials,omitzero"`
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
		return &env, err
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		return &env, err
	}
	if err := checkVaultConfig(env.Vault); err != nil {
		return &env, err
	}
	if err := checkAgentType(env.Agent); err != nil {
		return &env, err
	}
	return &env, nil
}

func NewAppConfig(env *Env) (*Env, error) {
	AppConfig = env
	return AppConfig, nil
}

/*
checkAgentType валидирует конфигурацию агента.
*/
func checkAgentType(agent agent) error {
	switch agent.Type {
	case "ssh":
		return nil
	case "portainer":
		if agent.Credentials.Url == "" {
			return fmt.Errorf("No parse portainer.url")
		}
		if agent.Credentials.Key == "" {
			return fmt.Errorf("No parse portainer.api_key")
		}
		return nil
	default:
		return errors.New("agent type is invalid")
	}
}

/*
checkVaultConfig валидирует конфигурацию Vault.
*/
func checkVaultConfig(vault vault) error {
	if vault.Url == "" {
		return fmt.Errorf("No parse vault.url")
	}
	if vault.Role == "" || vault.Secret == "" {
		return fmt.Errorf("No parse vault.role_id or vault.secret_id")
	}
	return nil
}
