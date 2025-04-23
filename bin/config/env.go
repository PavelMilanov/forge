package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Env описывает конфигурацию приложения.
type Env struct {
	Vault     vault
	Registry  registry
}

// server описывает конфигурацию сервера.
type vault struct {
	Url   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
	Path  string `mapstructure:"path"`
	Mount string `mapstructure:"mount"`
}

// registry описывает конфигурацию хранилища.
type registry struct {
	Url      string `mapstructure:"url"`
	Login    string `mapstructure:"login"`
	Password string `mapstructure:"password"`
}

func NewEnv(path, file string) *Env {
	env := Env{}
	viper.SetConfigName(file) // имя файла без расширения
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("не найден файл конфигурации : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		fmt.Println("не загружен файл конфигурации: ", err)
	}
	return &env
}
