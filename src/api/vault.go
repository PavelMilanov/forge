package api

import (
	"context"

	"github.com/PavelMilanov/forge/config"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/api/auth/approle"
)

/*
VaultAPI предоставляет интерфейс для взаимодействия с Vault.
*/
type Vault struct {
	ENV    *config.Env
	Client *api.Client
	API    *api.KVv2
}

/*
NewVaultClient обёртка над API Vault.

Returns

	*Vault - экземпляр Vault
	error - ошибка при создании экземпляра Vault
*/
func NewVaultClient(env *config.Env) (*Vault, error) {
	// env, err := config.NewEnv(config.FORGE_PATH, config.FORGE_FILE)
	// if err != nil {
	// 	return nil, err
	// }
	cfg := api.DefaultConfig()
	cfg.Address = env.Vault.Url
	client, err := api.NewClient(cfg)
	client.KVv2(config.VAULT_PATH)
	if err != nil {
		return nil, err
	}
	return &Vault{
		ENV:    env,
		Client: client,
		API:    client.KVv2(config.VAULT_PATH),
	}, nil
}

/*
Login авторизует приложение в Vault.

Returns

	error - ошибка при авторизации приложения в Vault
*/
func (v *Vault) Login() error {
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
	v.Client.SetToken(authInfo.Auth.ClientToken)
	// fmt.Printf("Успешная авторизация. Токен %s", authInfo.Auth.ClientToken)
	return nil
}
