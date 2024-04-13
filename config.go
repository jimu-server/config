package config

import (
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
		Redis    struct {
			Host     string
			Port     int
			Password string
			DB       int
		}
		Key    string
		OpenIm struct {
			Secret string
			Admin  string
			Init   struct {
				PlatformID           int32
				ApiAddr              string
				WsAddr               string
				DataDir              string
				LogLevel             uint32
				IsLogStandardOutput  bool
				LogFilePath          string
				IsExternalExtensions bool
			}
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

// Environment
// description: 加载配置
func environment() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(Evn); err != nil {
		panic(err)
	}
	return
}
