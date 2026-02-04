package config

import "path/filepath"

var (
	VERSION       = "dev"
	FORGE_PATH    = "var/forge"
	TEMPLATE_PATH = filepath.Join(FORGE_PATH, "templates")
	CONFIG_PATH   = filepath.Join(FORGE_PATH, "config")
	VAULT_PATH    = "forge"
	FORGE_FILE    = "forge.yml"
)

var SPECMODE = map[string]string{
	"swarm":      "swarm",
	"compose":    "compose",
	"kubernetes": "kubernetes",
}

var COMPOSEPARAMS = [2]string{
	"image",
	"tag",
}

var SWARMPARAMS = [3]string{
	"image",
	"tag",
	"replicas",
}

var KUBERNETESPARAMS = [2]string{
	"image",
	"tag",
}
