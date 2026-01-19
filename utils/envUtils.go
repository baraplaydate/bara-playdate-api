package utils

import (
	"bara-playdate-api/constant"
	"bara-playdate-api/model"
	"log"
	"time"
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

func NewEnv(cfg *model.ConsulConfigReq) Config {

	configFromConsul, err := ExportConfigFromConsul(cfg)
	if err != nil {
		log.Fatalf("Error loading config from Consul: %v", err)
	}

	// keyValues, err := ExportConfigFromConsul()
	// if err != nil {
	// 	log.Fatalf("Failed to get all keys and values: %v", err)
	// }

	// configFromConsul, err := UnmarshalConfig(keyValues)
	// if err != nil {
	// 	log.Fatalf("Failed to unmarshal config: %v", err)
	// }

	var config Config

	config.AppName = configFromConsul.AppName
	config.ServerPort = configFromConsul.ServerPort
	config.Environment = configFromConsul.Environment

	config.LogDirectory = configFromConsul.LogDirectory + " bara-playdate-api " + time.Now().Format("02-Jan-2006") + ".log"

	dbConnection, _ := DecryptAes256Sha256([]byte(configFromConsul.DbConnection), constant.KEY_AES)
	config.DbConnection = string(dbConnection)
	dbSchema, _ := DecryptAes256Sha256([]byte(configFromConsul.DbSchema), constant.KEY_AES)
	config.DbSchema = string(dbSchema)
	dbUsername, _ := DecryptAes256Sha256([]byte(configFromConsul.DbUsername), constant.KEY_AES)
	config.DbUsername = string(dbUsername)
	dbPassword, _ := DecryptAes256Sha256([]byte(configFromConsul.DbPassword), constant.KEY_AES)
	config.DbPassword = string(dbPassword)
	dbUrl, _ := DecryptAes256Sha256([]byte(configFromConsul.DbUrl), constant.KEY_AES)
	config.DbUrl = string(dbUrl)
	dbPort, _ := DecryptAes256Sha256([]byte(configFromConsul.DbPort), constant.KEY_AES)
	config.DbPort = string(dbPort)
	dbMaxPoolOpen, _ := DecryptAes256Sha256([]byte(configFromConsul.DbMaxPoolOpen), constant.KEY_AES)
	config.DbMaxPoolOpen = string(dbMaxPoolOpen)
	dbMaxPoolIdle, _ := DecryptAes256Sha256([]byte(configFromConsul.DbMaxPoolIdle), constant.KEY_AES)
	config.DbMaxPoolIdle = string(dbMaxPoolIdle)
	dbMaxPollLifeTime, _ := DecryptAes256Sha256([]byte(configFromConsul.DbMaxPollLifeTime), constant.KEY_AES)
	config.DbMaxPollLifeTime = string(dbMaxPollLifeTime)

	redisHost, _ := DecryptAes256Sha256([]byte(configFromConsul.RedisHost), constant.KEY_AES)
	config.RedisHost = string(redisHost)
	redisPort, _ := DecryptAes256Sha256([]byte(configFromConsul.RedisPort), constant.KEY_AES)
	config.RedisPort = string(redisPort)
	redisMaxSize, _ := DecryptAes256Sha256([]byte(configFromConsul.RedisMaxSize), constant.KEY_AES)
	config.RedisMaxSize = string(redisMaxSize)
	redisMinIdleSize, _ := DecryptAes256Sha256([]byte(configFromConsul.RedisMinIdleSize), constant.KEY_AES)
	config.RedisMinIdleSize = string(redisMinIdleSize)

	config.ApiKey = configFromConsul.ApiKey
	config.ApiKeyEncode = configFromConsul.ApiKeyEncode
	config.SignatureKey = configFromConsul.SignatureKey
	config.SignatureKeyEncode = configFromConsul.SignatureKeyEncode

	config.Route = configFromConsul.Route

	return config
}
