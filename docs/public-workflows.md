# Публичные GitHub Actions

`forge` можно использовать как публичный composite action из других репозиториев.

## Composite action

Путь в этом репозитории:

- `.github/actions/forge-set-tag/action.yml`

Что делает action:

- использует существующий `var/forge/forge.yml` на runner;
- требует входной параметр `config_output_path` с расширением `.yml` или `.yaml`;
- обновляет tag проекта в Vault через `forge env set <project> -p tag=<short_sha>`;
- рендерит конфигурацию через `forge env get <project> -c`.

Примечание:
- текущий action покрывает только операции `env`;
- деплой стека через Portainer нужно выполнять отдельным job, например:
  - `forge deploy stack file <endpoint> -n <stack-name> -f <stack.yml> --mode upsert`.

## Требования к runner

- бинарник `forge` установлен и доступен в `PATH`;
- файл `var/forge/forge.yml` существует и содержит валидные учетные данные Vault.

## Пример использования из другого репозитория

```yaml
name: Stage Build

on:
  push:
    branches: [stage]

jobs:
  forge-vars:
    runs-on: self-hosted
    outputs:
      tag: ${{ steps.forge_vars.outputs.tag }}
      config_path: ${{ steps.forge_vars.outputs.config_path }}
    steps:
      - name: Checkout
        uses: actions/checkout@v5

      - name: Set forge vars
        id: forge_vars
        uses: https://github.com/PavelMilanov/forge/.github/actions/forge-set-tag@main
        with:
          project: stage
          config_output_path: var/forge/stage.yml

  deploy:
    runs-on: self-hosted
    needs: forge-vars
    env:
      TAG: ${{ needs.forge-vars.outputs.tag }}
      CONFIG_PATH: ${{ needs.forge-vars.outputs.config_path }}
    steps:
      - name: Deploy
        run: |
          docker --context=stage stack deploy -c "$CONFIG_PATH" --with-registry-auth --detach=false admin
```

## Примечания

- Держите репозиторий `forge` публичным, чтобы внешние репозитории могли подключать action через `uses`.
- Перед запуском action проверьте, что `forge` доступен в `PATH` на self-hosted runner.
- Перед запуском action проверьте, что файл `var/forge/forge.yml` уже создан.
- Всегда передавайте `config_output_path` с расширением `.yml` или `.yaml`.
