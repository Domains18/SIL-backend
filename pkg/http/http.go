package http

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./conf/docker-config"
	}
	return "./conf/configs"
}