package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config는 애플리케이션 전체 설정을 담는 구조체
var Config AppConfig

// AppConfig는 애플리케이션 설정 구조체
type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
	Mongo    MongoConfig
	JWT      JWTConfig
}

// ServerConfig는 서버 관련 설정
type ServerConfig struct {
	Port string
}

// DatabaseConfig는 데이터베이스 관련 설정
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// MongoConfig는 MongoDB 관련 설정
type MongoConfig struct {
	URI string
	DB  string
}

// JWTConfig는 JWT 관련 설정
type JWTConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessDuration     time.Duration
	RefreshDuration    time.Duration
}

// 설정 파일 초기화
func Init() {
	// Viper 설정
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("../../configs")
	viper.AddConfigPath("../../../configs")

	// 환경 변수 대체 활성화
	viper.AutomaticEnv()

	// 설정 파일 읽기
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("설정 파일 읽기 오류: %v", err)
	}

	// 서버 설정
	Config.Server.Port = viper.GetString("server_port")

	// 데이터베이스 설정
	Config.Database.Host = viper.GetString("db_host")
	Config.Database.Port = viper.GetString("db_port")
	Config.Database.User = viper.GetString("db_user")
	Config.Database.Password = viper.GetString("db_password")
	Config.Database.DBName = viper.GetString("db_name")
	Config.Database.SSLMode = viper.GetString("db_sslmode")

	// MongoDB 설정
	Config.Mongo.URI = viper.GetString("mongo_uri")
	Config.Mongo.DB = viper.GetString("mongo_db")

	// JWT 설정
	Config.JWT.AccessTokenSecret = viper.GetString("jwt_access_token_secret")
	Config.JWT.RefreshTokenSecret = viper.GetString("jwt_refresh_token_secret")
	
	// 문자열로 된 기간 값을 Duration으로 변환
	accessDurationStr := viper.GetString("jwt_access_duration")
	if accessDurationStr != "" {
		if duration, err := time.ParseDuration(accessDurationStr); err == nil {
			Config.JWT.AccessDuration = duration
		}
	}
	
	refreshDurationStr := viper.GetString("jwt_refresh_duration")
	if refreshDurationStr != "" {
		if duration, err := time.ParseDuration(refreshDurationStr); err == nil {
			Config.JWT.RefreshDuration = duration
		}
	}

	// 중요한 설정 값에 대한 유효성 검사
	validateConfig()
}

// 설정 유효성 검사
func validateConfig() {
	// JWT 시크릿 키가 비어있는지 확인
	if Config.JWT.AccessTokenSecret == "" {
		log.Fatal("JWT Access Token Secret이 설정되지 않았습니다")
	}
	if Config.JWT.RefreshTokenSecret == "" {
		log.Fatal("JWT Refresh Token Secret이 설정되지 않았습니다")
	}
	// JWT 토큰 기간이 설정되지 않았으면 기본값 설정
	if Config.JWT.AccessDuration == 0 {
		Config.JWT.AccessDuration = 15 * time.Minute
	}
	if Config.JWT.RefreshDuration == 0 {
		Config.JWT.RefreshDuration = 7 * 24 * time.Hour // 7일
	}

	// 데이터베이스 설정 확인
	if Config.Database.Host == "" || Config.Database.Port == "" || Config.Database.User == "" ||
		Config.Database.Password == "" || Config.Database.DBName == "" {
		log.Fatal("postgresql 데이터베이스 설정이 완전하지 않습니다")
	}

	// MongoDB 설정 확인
	if Config.Mongo.URI == "" || Config.Mongo.DB == "" {
		log.Fatal("mongodb 설정이 완전하지 않습니다")
	}
}
