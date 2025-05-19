package utils

import (
	"fmt"

	"github.com/PavelMilanov/forge/config"
	"github.com/hashicorp/vault/api"
)

func NewVault() *api.KVv2 {
	env := config.NewEnv(config.CONFIG_PATH, "forge.yml")
	config := api.DefaultConfig()
	config.Address = env.Vault.Url
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Printf("Ошибка создания Vault клиента: %v", err)
	}
	client.SetToken(env.Vault.Token)
	kv := client.KVv2(env.Vault.Path)
	return kv
}
