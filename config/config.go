package config

import (
	"os"
)

func LoadingYaml() {
	// 读取YAML文件内容
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	server()
}
