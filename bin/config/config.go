package config

var (
	DURATION    = 2
	VERSION     = "v0.1.0"
	CONFIG_PATH = "var/config/"
	VAULT_PATH  = "forge"
)

var DOCKERMOD = map[string]int{
	"stack":   0,
	"compose": 1,
}
