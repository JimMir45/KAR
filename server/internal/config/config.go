package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Database  DatabaseConfig  `mapstructure:"database"`
	JWT       JWTConfig       `mapstructure:"jwt"`
	WeChat    WeChatConfig    `mapstructure:"wechat"`
	WeChatPay WeChatPayConfig `mapstructure:"wechat_pay"`
	COS       COSConfig       `mapstructure:"cos"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	SSLMode      string `mapstructure:"sslmode"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

type JWTConfig struct {
	Secret      string `mapstructure:"secret"`
	ExpireHours int    `mapstructure:"expire_hours"`
}

type WeChatConfig struct {
	Client WeChatAppConfig `mapstructure:"client"`
	Admin  WeChatAppConfig `mapstructure:"admin"`
}

type WeChatAppConfig struct {
	AppID     string `mapstructure:"app_id"`
	AppSecret string `mapstructure:"app_secret"`
}

type WeChatPayConfig struct {
	MchID            string `mapstructure:"mch_id"`
	MchAPIV3Key      string `mapstructure:"mch_api_v3_key"`
	MchSerialNumber  string `mapstructure:"mch_serial_number"`
	MchPrivateKeyPath string `mapstructure:"mch_private_key_path"`
	NotifyURL        string `mapstructure:"notify_url"`
}

type COSConfig struct {
	SecretID  string `mapstructure:"secret_id"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return cfg, nil
}
