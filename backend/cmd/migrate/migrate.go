package main

import (
	"fmt"
	"log"

	"time"

	"github.com/kamva/mgm/v3"
	channelModel "github.com/shjk0531/moye/backend/internal/domain/lounge/channel/model"
	loungeModel "github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/model"
	notificationModel "github.com/shjk0531/moye/backend/internal/domain/notification/model"
	recruitmentModel "github.com/shjk0531/moye/backend/internal/domain/recruitment/model"
	userModel "github.com/shjk0531/moye/backend/internal/domain/user/model"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/shjk0531/moye/backend/internal/global/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 설정 로드
	config.Init()
	
	// DSN 생성 (SSL 모드는 개발환경에 맞게 disable)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		config.Config.Database.Host,
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.DBName,
		config.Config.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %v", err)
	}

	// UUID 확장 활성화
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Fatalf("UUID 확장 활성화 실패: %v", err)
	}

	// 각 도메인 모델의 마이그레이션 실행
	err = db.AutoMigrate(
		&userModel.User{},
		&loungeModel.Lounge{},
		&loungeModel.LoungeMember{},
		&loungeModel.LoungeMemberRole{},
		&recruitmentModel.Recruitment{},
		&channelModel.Channel{},
		&channelModel.ChannelOrder{},
		&channelModel.DMChannel{},
		&channelModel.ChannelGroup{},
		&channelModel.ChannelGroupOrder{},
		&channelModel.VideoChatSession{},
		&notificationModel.Notification{},
	)
	
	if err != nil {
		log.Fatalf("마이그레이션 실패: %v", err)
	}

	log.Println("PostgreSQL 마이그레이션 성공!")

	// MongoDB 초기화 (mgm)
	initMongo()  // 위에서 정의한 초기화 함수 호출
	log.Println("MongoDB 연결 성공!")

	// (선택) 예시 데이터 삽입: 간단한 시간 값 테스트
	log.Println("현재 시간:", time.Now())
}

func initMongo() {
    err := mgm.SetDefaultConfig(nil, config.Config.Mongo.DB, options.Client().ApplyURI(config.Config.Mongo.URI))
    if err != nil {
        log.Fatalf("MongoDB 연결 실패: %v", err)
    }
}