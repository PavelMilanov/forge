# Public Action: forge-deploy-swarm

Путь action в репозитории:
- `.github/actions/forge-deploy-swarm/action.yml`

## Назначение

`forge-deploy-swarm` деплоит Docker Swarm stack через Portainer API:
- если stack не существует, создает его;
- если stack существует, обновляет его;
- может назначить права управления stack для команд Portainer.
- резолвит `endpoint_id` по `endpoint_name` через команду `forge deploy list`.

## Inputs

- `endpoint_name` (required): имя endpoint в Portainer.
- `stack_name` (required): имя стека.
- `stack_file` (required): путь к yaml-файлу стека.
- `prune` (optional, default `"true"`): удалять неиспользуемые сервисы при update.
- `pull_image` (optional, default `"false"`): принудительно подтягивать образы при update.
- `team_ids` (optional, default `""`): CSV список team id в Portainer (например `5` или `5,7`).

## Outputs

- `stack_id`: id стека в Portainer.
- `action`: выполненное действие (`created` или `updated`).

## Переменные окружения

Action не принимает `portainer_url` и `portainer_token` через `with`.
Нужно передать через `env`:

- `PORTAINER_URL`
- `PORTAINER_TOKEN`

Также на runner должен быть настроен `forge` (чтобы `forge deploy list` возвращал список endpoint-ов).

## Пример

```yaml
jobs:
  deploy:
    runs-on: self-hosted
    env:
      PORTAINER_URL: ${{ vars.PORTAINER_URL }}
      PORTAINER_TOKEN: ${{ secrets.PORTAINER_TOKEN }}
    steps:
      - uses: actions/checkout@v5

      - id: swarm
        uses: https://github.com/PavelMilanov/forge/.github/actions/forge-deploy-swarm@main
        with:
          endpoint_name: stage
          stack_name: admin
          stack_file: var/forge/stage.yml
          team_ids: "5,7"

      - run: |
          echo "Action: ${{ steps.swarm.outputs.action }}"
          echo "Stack ID: ${{ steps.swarm.outputs.stack_id }}"
```
