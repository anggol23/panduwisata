package config

import (
	"context"
	"mini-project/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBLog *mongo.Client

func InitDB() {
	dsn := "root:anggol2332prada@tcp(127.0.0.1:3306)/pandu-wisata?charset-utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitLog() {
	var err error
	DBLog, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/pandu-wisata"))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err = DBLog.Connect(ctx)
	if err != nil {
		panic(err)
	}

	DBLog.ListDatabaseNames(ctx, bson.M{})
}

func InitMigration() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.TempatWisata{})
	DB.AutoMigrate(&model.Transaction{})
	DB.AutoMigrate(&model.Category{})

}
