version=

dev:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -trimpath \
	-ldflags="-s -w -X 'github.com/PavelMilanov/forge/config.CONFIG_PATH=/var/forge'" \
	-o ../vagrant/server/work \
	.

prod:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go install -trimpath \
	-ldflags="-s -w -X 'github.com/PavelMilanov/forge/config.CONFIG_PATH=/var/forge' -X 'github.com/PavelMilanov/forge/config.VERSION=v${version}'"
