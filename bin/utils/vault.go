package utils

import (
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
func NewVaultClient() (*VaultClient, error) {
	env, err := config.NewEnv(config.CONFIG_PATH, config.CONFIG_FILE)
	if err != nil {
		return nil, err
	}
	config := api.DefaultConfig()
	config.Address = env.Vault.Url
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	client.SetToken(env.Vault.Token)
	_, err = client.Auth().Token().RenewSelf(2764800) // 30 days
	if err != nil {
		return nil, err
	}

	kv := client.KVv2(env.Vault.Path)
	return &VaultClient{
		ENV: env,
		KV:  kv,
	}, nil
}
