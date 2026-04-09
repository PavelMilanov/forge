# Public Action: forge-set-tag

Путь action в репозитории:
- `.github/actions/forge-set-tag/action.yml`

## Назначение

`forge-set-tag` обновляет tag проекта в Vault и рендерит итоговый yaml-конфиг.

Action:
- использует установленный на runner бинарник `forge`;
- выполняет `forge env set <project> -e <environment> -p api.image=<image>:<short_sha>`;
- выполняет `forge env get <project> -e <environment> -c > <config_output_path>`.

## Inputs

- `project` (required): имя проекта в Vault, например `admin-stage`.
- `environment` (required): имя окружения, например `stage`.
- `config_output_path` (required): путь выходного файла с расширением `.yml` или `.yaml`.

## Outputs

- `tag`: короткий git sha.
- `config_path`: путь к сгенерированному yaml-файлу.

## Требования

- runner: `self-hosted`;
- `forge` доступен в `PATH`;
- `var/forge/forge.yml` уже настроен на runner.

## Пример

```yaml
jobs:
  forge-vars:
    runs-on: self-hosted
    outputs:
      tag: ${{ steps.forge_vars.outputs.tag }}
      config_path: ${{ steps.forge_vars.outputs.config_path }}
    steps:
      - uses: actions/checkout@v5

      - id: forge_vars
        uses: https://github.com/PavelMilanov/forge/.github/actions/forge-set-tag@main
        with:
          project: api
          environment: stage
          config_output_path: var/forge/stage.yml
```
