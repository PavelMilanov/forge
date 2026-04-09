# Публичные GitHub Actions

`forge` можно использовать как публичный composite action из других репозиториев.

## Composite actions

Пути в этом репозитории:

- `.github/actions/forge-set-tag/action.yml`
- `.github/actions/forge-deploy-swarm/action.yml`

Что делает `forge-set-tag`:

- использует существующий `var/forge/forge.yml` на runner;
- требует входной параметр `config_output_path` с расширением `.yml` или `.yaml`;
- обновляет tag проекта в Vault через `forge env set <project> -p tag=<short_sha>`;
- рендерит конфигурацию через `forge env get <project> -c`.

Что делает `forge-deploy-swarm`:
- создает stack в Portainer, если его еще нет;
- обновляет stack, если он уже существует;
- работает с compose-файлом (`stack_file`) через Portainer API.
- поддерживает назначение прав на управление стеком для команд Portainer через `team_ids`.

Важно:
- `forge-deploy-swarm` не принимает `portainer_url` и `portainer_token` через `with`;
- action читает их только из переменных окружения:
  - `PORTAINER_URL`
  - `PORTAINER_TOKEN`

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
      PORTAINER_URL: ${{ vars.PORTAINER_URL }}
      PORTAINER_TOKEN: ${{ secrets.PORTAINER_TOKEN }}
    steps:
      - name: Checkout
        uses: actions/checkout@v5

      - name: Deploy swarm stack via Portainer
        id: swarm
        uses: https://github.com/PavelMilanov/forge/.github/actions/forge-deploy-swarm@main
        with:
          endpoint_name: stage
          stack_name: admin
          stack_file: ${{ env.CONFIG_PATH }}
          team_ids: "5,7"

      - name: Deployment summary
        run: |
          echo "Action: ${{ steps.swarm.outputs.action }}"
          echo "Stack ID: ${{ steps.swarm.outputs.stack_id }}"
```

## Примечания

- Держите репозиторий `forge` публичным, чтобы внешние репозитории могли подключать action через `uses`.
- Перед запуском action проверьте, что `forge` доступен в `PATH` на self-hosted runner.
- Перед запуском action проверьте, что файл `var/forge/forge.yml` уже создан.
- Всегда передавайте `config_output_path` с расширением `.yml` или `.yaml`.
- Для `forge-deploy-swarm` обязательно задайте `PORTAINER_URL` и `PORTAINER_TOKEN` через `env`.
- `team_ids` в `forge-deploy-swarm` передается как CSV (например, `5` или `5,7`).
