package api

import (
	"github.com/PavelMilanov/forge/config"
	"github.com/hashicorp/vault/api"
)

/*
VaultAPI предоставляет интерфейс для взаимодействия с Vault.
*/
type VaultAPI struct {
	ENV    *config.Env
	Client *api.Client
	API    *api.KVv2
}

/*
NewVaultClient обёртка над API Vault.

Returns

	*VaultAPI - экземпляр VaultAPI
	error - ошибка при создании экземпляра VaultAPI
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
	return &VaultAPI{
		ENV:    env,
		Client: client,
	}, nil
}

/*
RenewToken обновляет токен Vault.

Returns

	error - ошибка при обновлении токена Vault
*/
func (v *VaultAPI) RenewToken() error {
	_, err := v.Client.Auth().Token().RenewSelf(2764800) // 30 days
	if err != nil {
		return err
	}
	return nil
}

/*
Set устанавливает токен Vault.
*/
func (v *VaultAPI) Set() {
	v.Client.SetToken(v.ENV.Vault.Token)
	v.API = v.Client.KVv2(v.ENV.Vault.Path)
}
