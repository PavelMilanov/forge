# Public Workflow: forge-get-env

Путь workflow в репозитории:
- `.github/workflows/forge-get-env.yml`

## Назначение

`forge-get-env` — переиспользуемый workflow (`workflow_call`) для чтения данных окружения через обычную команду:
- `forge env get <project> -e <environment>`.

## Inputs

- `project` (required): имя проекта в Vault, например `api`.
- `environment` (required): имя окружения, например `stage`.

## Outputs

- `content`: содержимое вывода команды `forge env get`.

## Требования

- runner: `self-hosted`;
- `forge` доступен в `PATH`;
- `var/forge/forge.yml` уже настроен на runner.

## Пример вызова

```yaml
jobs:
  get_env_stage:
    uses: PavelMilanov/forge/.github/workflows/forge-get-env.yml@main
    with:
      project: api
      environment: stage

  use_output:
    runs-on: self-hosted
    needs: get_env_stage
    steps:
      - name: Show content size
        run: echo "${{ needs.get_env_stage.outputs.content }}" | wc -c
```
