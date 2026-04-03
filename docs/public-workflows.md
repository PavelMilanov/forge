# Public GitHub Workflows

`forge` can be used as a public reusable workflow from other repositories.

## Reusable workflow

Path in this repository:

- `.github/workflows/forge-set-tag.yml`

This workflow:

- installs `forge` binary;
- prepares `/var/forge/forge.yml` from repository secrets;
- sets tag in Vault using `forge env set <project> -p tag=<short_sha>`;
- renders deployment config using `forge env get <project> -c`.

Note:
- current reusable workflow covers only `env` operations;
- stack deployment via Portainer should be executed in a separate job using:
  - `forge deploy stack file <endpoint> -n <stack-name> -f <stack.yml> --mode upsert`.

## Required secrets in caller repo

- `VAULT_URL`
- `VAULT_ROLE_ID`
- `VAULT_SECRET_ID`

## Example usage from another repository

```yaml
name: Stage Build

on:
  push:
    branches: [stage]

jobs:
  forge-vars:
    uses: https://github.com/PavelMilanov/forge/.github/workflows/forge-set-tag.yml@v0.1.6
    with:
      project: stage
      # optional: config_output_path: /var/forge/stage.yml
      # optional: forge_download_url: https://github.com/PavelMilanov/forge/releases/download/v0.1.6/forge
    secrets:
      VAULT_URL: ${{ secrets.VAULT_URL }}
      VAULT_ROLE_ID: ${{ secrets.VAULT_ROLE_ID }}
      VAULT_SECRET_ID: ${{ secrets.VAULT_SECRET_ID }}

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

## Notes

- Keep `forge` repository public so external repositories can reference workflow via `uses`.
- Use pinned tags in both `uses:` and `forge_download_url` for deterministic runs.
