/**
 * 基础配置
 * 读取yaml文件, 解析配置
 */

package yaml

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Server 服务器配置
type Server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

// Database 数据库配置
type Database struct {
	Dialect  string `yaml:"dialect"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxOpen  int    `yaml:"max_open"`
}

// Redis 配置
type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"db"`
}

// Log 日志配置
type Log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

// Upload上传配置
type Upload struct {
	Host      string `yaml:"host"`
	ImagePath string `yaml:"imagePath"`
	FilePath  string `yaml:"filePath"`
	MaxSize   int    `yaml:"maxSize"`
	MaxNum    int    `yaml:"maxNum"`
	AllowExt  string `yaml:"allowExt"`
}

// 其他配置, 如密钥等
type App struct {
	JwtSecret      string `yaml:"jwtSecret"`
	JwtTokenExpire int    `yaml:"jwtTokenExpire"`
}

// 合并配置
type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"db"`
	Redis    Redis    `yaml:"redis"`
	Log      Log      `yaml:"log"`
	Upload   Upload   `yaml:"upload"`
	App      App      `yaml:"app"`
}

var Cfg *Config

func init() {
	yamlFile, err := os.ReadFile("./config.yaml")
	// 有错就宕机
	if err != nil {
		panic(err)
	}
	// 解析yaml文件, 绑定值
	err = yaml.Unmarshal(yamlFile, &Cfg)
	if err != nil {
		panic(err)
	}
}
