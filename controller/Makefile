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

token:
	docker exec -e VAULT_TOKEN='root' vault vault token create --renewable=true -policy=forge
