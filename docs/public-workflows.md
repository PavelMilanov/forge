# Публичные GitHub Actions

`forge` можно использовать как публичный composite action из других репозиториев.

Эта страница содержит навигацию. Подробные описания вынесены в отдельные независимые документы:

- [Public Action: forge-set-tag](public-action-forge-set-tag.md)
- [Public Action: forge-deploy-swarm](public-action-forge-deploy-swarm.md)

## Общие требования к runner

- бинарник `forge` установлен и доступен в `PATH`;
- файл `var/forge/forge.yml` существует и содержит валидные учетные данные Vault.
