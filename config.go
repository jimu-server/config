package config

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
)

var Evn = &Configuration{}

// Configuration 配置文件映射
type Configuration struct {
	App struct {
		Number int64
		Port   string
		Logger struct {
			Level      string
			FileName   string
			MaxSize    int
			MaxBackups int
			MaxAge     int
		}
		Database string
		FileDb   string
		Redis    struct {
			Host     string
			Port     int
			Password string
			DB       int
		}
		Key    string
		Ollama struct {
			Port string
		}
		Minio struct {
			Endpoint        string
			AccessKeyID     string
			SecretAccessKey string
			Secure          bool
		}
		Tencent struct {
			BucketURL  string
			ServiceURL string
			SecretID   string
			SecretKey  string
		}
		RabbitMQ struct {
			Host     string
			Port     int
			User     string
			Password string
		}
		Pay struct {
			AliPay struct {
				AppId           string
				AppPublicCert   string
				AliPayPublicKey string
				AliPayRootKey   string
			}
		}
		Email struct {
			Host     string
			Port     int
			User     string
			Password string
		}
	}
}

func init() {
	environment()
}

// 用于打包需要读取的配置文件
//
//go:embed  config.yml
var data []byte

// Environment
// description: 加载配置
func environment() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if err = viper.ReadConfig(bytes.NewBuffer(data)); err != nil {
			panic(err)
		}
	}
	if err := viper.Unmarshal(Evn); err != nil {
		panic(err)
	}
	return
}
