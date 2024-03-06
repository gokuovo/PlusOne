package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// Config 结构体用于映射YAML配置文件的结构
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	// 添加其他配置字段...
}

// ServerConfig 结构体用于映射服务器配置
type ServerConfig struct {
	Port int64 `yaml:"port"`
	// 添加其他服务器配置字段...
}

// DatabaseConfig 结构体用于映射数据库配置
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	// 添加其他数据库配置字段...
}

func LoadingYaml() {
	//获取当前包的路径
	_, currentFile, _, _ := runtime.Caller(0)
	packagePath := path.Dir(currentFile)

	// 构建YAML文件的路径
	yamlFilePath := filepath.Join(packagePath, "config.yaml")

	// 读取YAML文件内容
	yamlFile, err := os.ReadFile(yamlFilePath)
	if err != nil {
		panic(err)
	}

	// 定义一个Config结构体变量
	var config Config

	// 解析YAML内容到Config结构体
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	//连接数据库
	mysqlCon(config.Database)
	//启动服务
	server(config.Server.Port)
}
