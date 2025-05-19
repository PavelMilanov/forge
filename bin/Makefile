dev:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -trimpath \
	-ldflags="-s -w -X 'github.com/PavelMilanov/forge/config.CONFIG_PATH=/var/forge'" \
	-o ../vagrant/server/work \
	.
