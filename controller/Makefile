version=

dev:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -trimpath \
	-ldflags="-s -w -X 'github.com/PavelMilanov/forge/config.FORGE_PATH=/var/forge'" \
	-o ../vagrant/server/work \
	.

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go install -trimpath \
	-ldflags="-s -w -X 'github.com/PavelMilanov/forge/config.FORGE_PATH=/var/forge' -X 'github.com/PavelMilanov/forge/config.VERSION=${version}'"

start:
	docker exec -e VAULT_TOKEN='root' vault vault auth enable approle
	#vault write auth/approle/role/forge token_period="24h" token_policies="my-policy"
	#vault read auth/approle/role/forge/role-id
	#vault write -f auth/approle/role/forge/secret-id
