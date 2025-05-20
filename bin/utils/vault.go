package utils

import (
	"fmt"

	"github.com/PavelMilanov/forge/config"
	"github.com/hashicorp/vault/api"
)

type VaultClient struct {
	ENV *config.Env
	KV  *api.KVv2
}

func NewVaultClient() *VaultClient {
	env := config.NewEnv(config.CONFIG_PATH, "forge.yml")
	config := api.DefaultConfig()
	config.Address = env.Vault.Url
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}
	client.SetToken(env.Vault.Token)
	kv := client.KVv2(env.Vault.Path)
	return &VaultClient{
		ENV: env,
		KV:  kv,
	}
}
