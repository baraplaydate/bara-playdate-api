package utils

import (
	"bara-playdate-api/exception"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config Config) *gorm.DB {
	username := config.DbUsername
	password := config.DbPassword
	host := config.DbUrl
	port := config.DbPort
	database := config.DbConnection
	maxPoolOpen, err := strconv.Atoi("10")
	maxPoolIdle, err := strconv.Atoi("5")
	maxPollLifeTime, err := strconv.Atoi("30000")
	exception.PanicLogging(err)

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: loggerDb,
	})
	exception.PanicLogging(err)

	sqlDB, err := db.DB()
	exception.PanicLogging(err)

	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	exception.PanicLogging(err)

	return db
}
