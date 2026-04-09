# Public Workflow: forge-set-env

Путь workflow в репозитории:
- `.github/workflows/forge-set-env.yml`

## Назначение

`forge-set-env` — переиспользуемый workflow (`workflow_call`) для обновления данных окружения через `forge env set`.

Workflow полностью совместим с CLI `forge env set`:
- каждый параметр из `params` передается как отдельный флаг `-p key=value`;
- поддерживаются вложенные ключи (`api.image`, `placement.enabled`, `api.replicas` и т.д.).

## Inputs

- `project` (required): имя проекта в Vault, например `api`.
- `environment` (required): имя окружения, например `stage`.
- `params` (required): многострочный список `key=value`, где каждая строка превращается в `-p key=value`.

## Outputs

- `applied_params`: количество примененных параметров `-p`.

## Требования

- runner: `self-hosted`;
- `forge` доступен в `PATH`;
- `var/forge/forge.yml` уже настроен на runner.

## Пример вызова

```yaml
jobs:
  set_env:
    uses: PavelMilanov/forge/.github/workflows/forge-set-env.yml@main
    with:
      project: api
      environment: stage
      params: |
        api.image=registry.service.uhvahta.ru/api/backend:${{ github.sha }}
        api.replicas=1
        placement.enabled=false
```
