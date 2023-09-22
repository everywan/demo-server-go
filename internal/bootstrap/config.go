package bootstrap

import (
	"os"

	"github.com/spf13/viper"
)

// 一般而言, 大项目建议使用配置中心的方式, 小项目可以使用配置文件方式.
type configComponent struct {
	cfg *Config
}

// GetConfig 根据传入的路径解析
func (b *Bootstrap) GetConfig(path string) *Config {
	if b.cfg == nil {
		if path != "" {
			viper.SetConfigFile(path)
		} else {
			home, err := os.UserHomeDir()
			handleInitError("config_get_home_dir", err)

			viper.AddConfigPath(home)
			viper.SetConfigName(".demo-server-go.yaml")
		}
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		handleInitError("config_read_file", err)

		b.cfg = &Config{}
		err = viper.Unmarshal(b.cfg)
		handleInitError("config_unmarshal", err)
	}
	return b.cfg
}

type (
	GrpcConfig struct {
		Port int `mapstructure:"port" yaml:"port"`
	}
	ServerConfig struct {
		Host            string   `mapstructure:"host" yaml:"host"`
		Port            uint     `mapstructure:"port" yaml:"port"`
		CorsAllowOrigin []string `mapstructure:"cors_allow_origin"`
		DefaultExpired  int64    `mapstructure:"default_expired"`
	}
	Config struct {
		// Logging log.Options      `mapstructure:"logger"`
		// DB     database.Options `mapstructure:"db" yaml:"db"`
		Grpc   GrpcConfig   `mapstructure:"grpc" yaml:"grpc"`
		Server ServerConfig `mapstructure:"server" yaml:"server"`
	}
)
