package config

import (
	"bytes"
	_ "embed"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
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
//go:embed  config-dev.yml
var dev []byte

//go:embed  config-dev.yml
var product []byte

//go:embed  config-dev.yml
var test []byte

// Environment
// description: 加载配置
func environment() {
	active := pflag.String("active", "dev", "Activation configuration")
	pflag.Parse()
	name := "config"
	if active == nil {
		name = strings.Join([]string{name, "dev"}, "-")
	} else {
		name = strings.Join([]string{name, *active}, "-")
	}
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		switch *active {
		case "dev":
			err = viper.ReadConfig(bytes.NewBuffer(dev))
		case "product":
			err = viper.ReadConfig(bytes.NewBuffer(product))
		case "test":
			err = viper.ReadConfig(bytes.NewBuffer(test))
		}
		if err != nil {
			panic(err)
		}
	}
	if err := viper.Unmarshal(Evn); err != nil {
		panic(err)
	}
	return
}
