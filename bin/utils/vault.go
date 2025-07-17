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

/*
 * Initialize a new Vault client
 *
 */
func NewVaultClient() *VaultClient {
	env, err := config.NewEnv(config.CONFIG_PATH, "forge.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config := api.DefaultConfig()
	config.Address = env.Vault.Url
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client.SetToken(env.Vault.Token)
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
