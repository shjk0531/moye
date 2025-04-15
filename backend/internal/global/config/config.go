package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerPort string `mapstructure:"server_port"`

	// PostgreSQL 설정
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`

	// MongoDB 설정
	MongoURI string `mapstructure:"mongo_uri"`
	MongoDB  string `mapstructure:"mongo_db"`

	// JWT 설정
	JWT JWTConfig
}

// JWTConfig는 JWT 관련 설정을 담는 구조체
type JWTConfig struct {
	SecretKey     string        `mapstructure:"jwt_secret_key"`
	TokenDuration time.Duration `mapstructure:"jwt_token_duration"`
}

var Config AppConfig

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")

	// 기본값 설정
	viper.SetDefault("jwt_secret_key", "your-secure-jwt-secret-key-at-least-32-bytes-long")
	viper.SetDefault("jwt_token_duration", 24*time.Hour)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("설정 파일 읽기 실패: %v", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("설정값 매핑 실패: %v", err)
	}

	// JWT 설정 초기화
	Config.JWT = JWTConfig{
		SecretKey:     viper.GetString("jwt_secret_key"),
		TokenDuration: viper.GetDuration("jwt_token_duration"),
	}
}
