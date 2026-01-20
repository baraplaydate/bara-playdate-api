package utils

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppName            string
	ServerPort         string
	Environment        string
	AppMode            string
	LogDirectory       string
	DbConnection       string
	DbUsername         string
	DbSchema           string
	DbPassword         string
	DbUrl              string
	DbPort             string
	DbSid              string
	DPass              string
	DbMaxPoolOpen      string
	DbMaxPoolIdle      string
	DbMaxPollLifeTime  string
	RedisHost          string
	RedisPort          string
	RedisMaxSize       string
	RedisMinIdleSize   string
	ApiKey             string
	ApiKeyEncode       string
	SignatureKey       string
	SignatureKeyEncode string
	Route              string
}

func NewEnv() Config {
	var config Config

	viper.SetConfigFile("yaml")
	viper.AddConfigPath("./")
	viper.SetConfigName("config")

	if errData := viper.ReadInConfig(); errData != nil {
		fmt.Printf("Error reading config file, %s", errData)
	}

	config.AppName = viper.GetString("appName")
	config.ServerPort = viper.GetString("serverPort")
	config.Environment = viper.GetString("environment")
	if viper.GetString("releaseMode") == "y" || viper.GetString("releaseMode") == "Y" {
		config.AppMode = "release"
	} else {
		config.AppMode = "debug"
	}
	config.LogDirectory = viper.GetString("logDirectory."+config.Environment+".path") + " go-gin-clean-architecture " + time.Now().Format("02-Jan-2006") + ".log"

	dbConnection := viper.GetString("database." + config.Environment + ".connection")
	config.DbConnection = string(dbConnection)
	dbSchema := viper.GetString("database." + config.Environment + ".schema")
	config.DbSchema = string(dbSchema)
	dbUsername := viper.GetString("database." + config.Environment + ".username")
	config.DbUsername = string(dbUsername)
	dbPassword := viper.GetString("database." + config.Environment + ".password")
	config.DbPassword = string(dbPassword)
	dbUrl := viper.GetString("database." + config.Environment + ".url")
	config.DbUrl = string(dbUrl)
	dbPort := viper.GetString("database." + config.Environment + ".port")
	config.DbPort = string(dbPort)

	dbMaxPoolOpen := viper.GetString("database." + config.Environment + ".maxPoolOpen")
	config.DbMaxPoolOpen = string(dbMaxPoolOpen)
	dbMaxPoolIdle := viper.GetString("database." + config.Environment + ".maxPoolIdle")
	config.DbMaxPoolIdle = string(dbMaxPoolIdle)
	dbMaxPollLifeTime := viper.GetString("database." + config.Environment + ".maxPollLifeTime")
	config.DbMaxPollLifeTime = string(dbMaxPollLifeTime)

	redisHost := viper.GetString("redis." + config.Environment + ".host")
	config.RedisHost = string(redisHost)
	redisPort := viper.GetString("redis." + config.Environment + ".port")
	config.RedisPort = string(redisPort)
	redisMaxSize := viper.GetString("redis." + config.Environment + ".maxSize")
	config.RedisMaxSize = string(redisMaxSize)
	redisMinIdleSize := viper.GetString("redis." + config.Environment + ".minIdleSize")
	config.RedisMinIdleSize = string(redisMinIdleSize)

	config.ApiKey = viper.GetString("key." + config.Environment + ".apiKey")
	config.ApiKeyEncode = viper.GetString("key." + config.Environment + ".apiKeyEncode")
	config.SignatureKey = viper.GetString("key." + config.Environment + ".signatureKey")
	config.SignatureKeyEncode = viper.GetString("key." + config.Environment + ".signatureKeyEncode")

	config.Route = viper.GetString("route." + config.Environment + ".name")

	return config
}
