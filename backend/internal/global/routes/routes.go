// internal/global/routes/routes.go
package routes

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/global/middleware"

	// PostgreSQL: User 도메인
	userController "github.com/shjk0531/moye/backend/internal/domain/user/controller"
	userRepository "github.com/shjk0531/moye/backend/internal/domain/user/repository"

	// PostgreSQL: Study 도메인
	studyController "github.com/shjk0531/moye/backend/internal/domain/study/controller"
	studyRepository "github.com/shjk0531/moye/backend/internal/domain/study/repository"

	// MongoDB: chat 도메인
	chatContainer "github.com/shjk0531/moye/backend/internal/domain/chat/message/controller"
	chatRepository "github.com/shjk0531/moye/backend/internal/domain/chat/message/repository"

	"github.com/shjk0531/moye/backend/internal/global/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine) {
	// PostgreSQL 연결 (정형 데이터: User, Study, Notification 등)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		config.Config.Database.Host,
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.DBName,
		config.Config.Database.Port,
	)
	pgDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("PostgreSQL 연결 실패: " + err.Error())
	}

	api := router.Group("/api")
	// 공통 라우트
	public := api.Group("")
	// 보호된 라우트 (JWT 인증 필요)
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())

	// User 도메인 라우트 등록 (Public)
	userRepo := userRepository.NewRepository(pgDB)
	userRootCtrl := userController.Init(userRepo)
	userRootCtrl.RegisterPublicRoutes(public)

	// User 프로필 조회 (Protected)
	userRootCtrl.RegisterProtectedRoutes(protected)

	// Study 도메인 라우트 등록 (인증 필요)
	studyRepo := studyRepository.NewRepository(pgDB)
	studyCtrl := studyController.Init(studyRepo, pgDB)
	studyCtrl.RegisterPublicRoutes(public)
	studyCtrl.RegisterPrivateRoutes(protected)

	// MongoDB 연결 (메시지 데이터: 채팅, 로그 등)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Config.Mongo.URI))
	if err != nil {
		panic("MongoDB 연결 실패: " + err.Error())
	}
	// 연결 확인
	if err := mongoClient.Ping(ctx, nil); err != nil {
		panic("MongoDB ping 실패: " + err.Error())
	}

	// chat 도메인 라우트 등록 (인증 필요)
	msgRepo := chatRepository.NewRepository(mongoClient, config.Config.Mongo.DB)
	msgCtrl := chatContainer.NewController(msgRepo)
	msgCtrl.RegisterRoutes(protected)
}
