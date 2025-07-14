package utils

import (
	"fmt"
	"os"

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
	_, err = client.Auth().Token().LookupSelf()
	if err != nil {
		fmt.Printf("Failed to lookup token: %v\n", err)
		os.Exit(1)
	}
	_, err = client.Auth().Token().RenewSelf(2764800)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	kv := client.KVv2(env.Vault.Path)
	return &VaultClient{
		ENV: env,
		KV:  kv,
	}
}
