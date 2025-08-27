package api

import (
	"github.com/PavelMilanov/forge/config"
	"github.com/hashicorp/vault/api"
)

type VaultAPI struct {
	ENV    *config.Env
	Client *api.Client
	API    *api.KVv2
}

/*
 * Initialize a new Vault client
 *
 */
func NewVaultClient() (*VaultAPI, error) {
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
	_, err = client.Auth().Token().RenewSelf(2764800) // 30 days
	if err != nil {
		return nil, err
	}
	return &VaultAPI{
		ENV:    env,
		Client: client,
	}, nil
}

func (v *VaultAPI) RenewToken() error {
	_, err := v.Client.Auth().Token().RenewSelf(2764800) // 30 days
	if err != nil {
		return err
	}
	return nil
}

func (v *VaultAPI) Set() {
	v.Client.SetToken(v.ENV.Vault.Token)
	v.API = v.Client.KVv2(v.ENV.Vault.Path)
}
