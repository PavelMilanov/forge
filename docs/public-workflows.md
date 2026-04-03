# Public GitHub Actions

`forge` can be used as a public composite action from other repositories.

## Composite action

Path in this repository:

- `.github/actions/forge-set-tag/action.yml`

This action:

- runs in the caller job as regular `steps` (not as a separate job);
- requires `self-hosted` runner with preinstalled `forge`;
- verifies `forge version` before executing commands;
- requires existing `var/forge/forge.yml` on runner (Vault settings are read from this file);
- requires `config_output_path` input with file extension (`.yml` or `.yaml`);
- sets tag in Vault using `forge env set <project> -p tag=<short_sha>`;
- renders deployment config using `forge env get <project> -c`.

Note:
- current action covers only `env` operations;
- stack deployment via Portainer should be executed in a separate job using:
  - `forge deploy stack file <endpoint> -n <stack-name> -f <stack.yml> --mode upsert`.

## Required preconditions on runner

- `forge` binary is installed and available in `PATH`;
- config file exists at `var/forge/forge.yml` with valid Vault credentials.

## Example usage from another repository

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
        uses: https://github.com/PavelMilanov/forge/.github/actions/forge-set-tag@v0.1.6
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

- Убедитесь, что бинарник `forge` установлен и доступен в `PATH` на self-hosted runner.
- Убедитесь, что перед запуском action существует файл `var/forge/forge.yml`.
- Всегда передавайте `config_output_path` с расширением `.yml` или `.yaml`.
