package config

var (
	VERSION     = "dev"
	CONFIG_PATH = "var/config/"
	VAULT_PATH  = "forge"
	CONFIG_FILE = "forge.yml"
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
