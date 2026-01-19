package api

import (
	"context"
	"fmt"

	"github.com/PavelMilanov/forge/config"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/api/auth/approle"
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
	cfg := api.DefaultConfig()
	cfg.Address = env.Vault.Url
	client, err := api.NewClient(cfg)
	client.KVv2(config.VAULT_PATH)
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
// func (v *VaultAPI) Set() {
// 	v.Client.SetToken(v.ENV.Vault.Token)
// 	v.API = v.Client.KVv2(config.VAULT_PATH)
// }

/*
Login авторизует приложение в Vault.

Returns

	error - ошибка при авторизации приложения в Vault
*/
func (v *VaultAPI) Login() error {
	auth, err := approle.NewAppRoleAuth(
		v.ENV.Vault.Role,
		&approle.SecretID{FromString: v.ENV.Vault.Secret},
	)
	if err != nil {
		return err
	}
	authInfo, err := v.Client.Auth().Login(context.Background(), auth)
	if err != nil {
		return err
	}
	fmt.Printf("Успешная авторизация. Токен %s", authInfo.Auth.ClientToken)
	return nil
}
