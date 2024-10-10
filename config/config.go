package config

// DefaultConfigs maps services (AWS, kubectl, etc.) to their default config paths
var DefaultConfigs = map[string]string{
	"aws":         "~/.aws",
	"kubectl":     "~/.kube",
	"digitalocean": "~/.config/doctl",
}

