package config

var (
	VERSION     = "dev"
	CONFIG_PATH = "var/config/"
	VAULT_PATH  = "forge"
)

var DOCKERMOD = map[string]int{
	"stack":   0,
	"compose": 1,
}
