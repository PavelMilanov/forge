package config

var (
	VERSION     = "dev"
	CONFIG_PATH = "var/config/"
	VAULT_PATH  = "forge"
)

var SPECMODE = map[string]string{
	"swarm":      "swarm",
	"compose":    "compose",
	"kubernetes": "kubernetes",
}
