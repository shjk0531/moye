package db

import (
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/shjk0531/moye/backend/internal/global/config"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo() {
	err := mgm.SetDefaultConfig(nil, config.Config.Mongo.DB, options.Client().ApplyURI(config.Config.Mongo.URI))
	if err != nil {
		log.Fatal("MongoDB(mgm) 초기화 실패: ", err)
	}
}
