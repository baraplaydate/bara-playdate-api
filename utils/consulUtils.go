package utils

import (
	"bara-playdate-api/constant"
	"bara-playdate-api/model"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"
)

// LoadConsulConfig memuat YAML ke dalam struktur Config
func LoadConsulConfig() (*model.ConsulConfigReq, error) {

	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	// Unmarshal YAML data
	var config model.ConsulConfigReq
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("error unmarshalling YAML: %v", err)
	}

	// Output the config to verify
	fmt.Printf("Config: %+v\n", config)

	return &config, nil
}

// ImportConfigToConsul mengimpor konfigurasi ke Consul
func ImportConfigToConsul(cfg *model.ConsulConfigReq) error {
	//prefix dibutuhkan untuk membedakan project yang saya gunakan

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return fmt.Errorf("error creating Consul client: %w", err)
	}

	// Ambil environment yang aktif
	env := cfg.Environment

	// Fungsi untuk mengimpor konfigurasi ke Consul tanpa prefix
	// importKey := func(key string, value interface{}) {
	// 	_, err := client.KV().Put(&api.KVPair{Key: key, Value: []byte(fmt.Sprintf("%v", value))}, nil)
	// 	if err != nil {
	// 		log.Printf("Error putting key '%s' into Consul: %v", key, err)
	// 	}
	// }

	// Fungsi untuk mengimpor konfigurasi ke Consul dengan prefix
	importKey := func(key string, value interface{}) {
		fullKey := fmt.Sprintf("%s/%s", constant.PREFIX_CONSUL, key)
		_, err := client.KV().Put(&api.KVPair{Key: fullKey, Value: []byte(fmt.Sprintf("%v", value))}, nil)
		if err != nil {
			log.Printf("Error putting key '%s' into Consul: %v", fullKey, err)
		}
	}

	// Mengimpor variabel-variabel utama
	importKey("appName", cfg.AppName)
	importKey("serverPort", cfg.ServerPort)
	importKey("environment", cfg.Environment)
	importKey("releaseMode", cfg.ReleaseMode)

	// Mengimpor konfigurasi log directory
	if logDir, ok := cfg.LogDirectory[env]; ok {
		importKey(fmt.Sprintf("logDirectory/%s/path", env), logDir.Path)
	}

	// Mengimpor konfigurasi Redis
	if redisConfig, ok := cfg.Redis[env]; ok {
		importKey(fmt.Sprintf("redis/%s/host", env), redisConfig.Host)
		importKey(fmt.Sprintf("redis/%s/port", env), redisConfig.Port)
		importKey(fmt.Sprintf("redis/%s/maxSize", env), redisConfig.MaxSize)
		importKey(fmt.Sprintf("redis/%s/minIdleSize", env), redisConfig.MinIdleSize)
	}

	// Mengimpor konfigurasi Database
	if dbConfig, ok := cfg.Database[env]; ok {
		importKey(fmt.Sprintf("database/%s/connection", env), dbConfig.Connection)
		importKey(fmt.Sprintf("database/%s/username", env), dbConfig.Username)
		importKey(fmt.Sprintf("database/%s/password", env), dbConfig.Password)
		importKey(fmt.Sprintf("database/%s/url", env), dbConfig.URL)
		importKey(fmt.Sprintf("database/%s/port", env), dbConfig.Port)
		importKey(fmt.Sprintf("database/%s/schema", env), dbConfig.Schema)
		importKey(fmt.Sprintf("database/%s/maxPoolOpen", env), dbConfig.MaxPoolOpen)
		importKey(fmt.Sprintf("database/%s/maxPoolIdle", env), dbConfig.MaxPoolIdle)
		importKey(fmt.Sprintf("database/%s/maxPollLifeTime", env), dbConfig.MaxPollLifeTime)
	}

	// Mengimpor konfigurasi Key Access
	if keyConfig, ok := cfg.Key[env]; ok {
		importKey(fmt.Sprintf("key/%s/apiKey", env), keyConfig.APIKey)
		importKey(fmt.Sprintf("key/%s/apiKeyEncode", env), keyConfig.APIKeyEncode)
		importKey(fmt.Sprintf("key/%s/signatureKey", env), keyConfig.SignatureKey)
		importKey(fmt.Sprintf("key/%s/signatureKeyEncode", env), keyConfig.SignatureKeyEncode)
	}

	// Mengimpor konfigurasi Route
	if routeConfig, ok := cfg.Route[env]; ok {
		importKey(fmt.Sprintf("route/%s/name", env), routeConfig.Name)
	}

	log.Println("Config imported to Consul successfully")
	return nil
}

// func ExportConfigFromConsul() (map[string]string, error) {
// 	client, err := api.NewClient(api.DefaultConfig())
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating Consul client: %w", err)
// 	}

// 	kv := client.KV()

// 	keys, _, err := kv.List(constant.PREFIX_CONSUL, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("error listing keys from Consul: %w", err)
// 	}

// 	if len(keys) == 0 {
// 		return nil, fmt.Errorf("no keys found with prefix '%s'", constant.PREFIX_CONSUL)
// 	}

// 	keyValues := make(map[string]string)

// 	for _, kvPair := range keys {
// 		key := kvPair.Key
// 		pair, _, err := kv.Get(key, nil)
// 		if err != nil {
// 			return nil, fmt.Errorf("error getting key '%s' from Consul: %w", key, err)
// 		}

// 		if pair != nil {
// 			keyValues[key] = string(pair.Value)
// 		}
// 	}

// 	return keyValues, nil
// }

// // UnmarshalConfig unmarshals Consul data into ConsulConfig struct
// func UnmarshalConfig(data map[string]string) (*model.ConsulConfig, error) {
// 	config := &model.ConsulConfig{
// 		LogDirectory: make(map[string]model.LogDirectory),
// 		Redis:        make(map[string]model.RedisConfig),
// 		Database:     make(map[string]model.DatabaseConfig),
// 		Key:          make(map[string]model.KeyAccess),
// 		Route:        make(map[string]model.Route),
// 	}

// 	for key, value := range data {
// 		// Extract section and sub-key from Consul key
// 		parts := strings.Split(key, "/")
// 		if len(parts) < 3 {
// 			continue
// 		}

// 		section := parts[1]
// 		subKey := parts[2]

// 		// Use `subKey` for identifying specific entries within the section
// 		switch section {
// 		case "appName":
// 			config.AppName = value
// 		case "serverPort":
// 			config.ServerPort = value
// 		case "environment":
// 			config.Environment = value
// 		case "releaseMode":
// 			config.ReleaseMode = value
// 		case "logDirectory":
// 			var logDir model.LogDirectory
// 			if err := json.Unmarshal([]byte(value), &logDir); err != nil {
// 				return nil, err
// 			}
// 			config.LogDirectory[subKey] = logDir
// 		case "redis":
// 			var redis model.RedisConfig
// 			if err := json.Unmarshal([]byte(value), &redis); err != nil {
// 				return nil, err
// 			}
// 			config.Redis[subKey] = redis
// 		case "database":
// 			var database model.DatabaseConfig
// 			if err := json.Unmarshal([]byte(value), &database); err != nil {
// 				return nil, err
// 			}
// 			config.Database[subKey] = database
// 		case "key":
// 			var keyAccess model.KeyAccess
// 			if err := json.Unmarshal([]byte(value), &keyAccess); err != nil {
// 				return nil, err
// 			}
// 			config.Key[subKey] = keyAccess
// 		case "route":
// 			var route model.Route
// 			if err := json.Unmarshal([]byte(value), &route); err != nil {
// 				return nil, err
// 			}
// 			config.Route[subKey] = route
// 		}
// 	}

// 	return config, nil
// }

// ExportConfigFromConsul memuat konfigurasi dari Consul dengan prefix tertentu
func ExportConfigFromConsul(cfg *model.ConsulConfigReq) (*model.ConsulConfigRes, error) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("error creating Consul client: %w", err)
	}

	// Ambil environment yang aktif
	env := cfg.Environment

	kv := client.KV()

	// Mengambil daftar kunci dengan prefix tertentu
	keys, _, err := kv.List(constant.PREFIX_CONSUL, nil)
	if err != nil {
		return nil, fmt.Errorf("error listing keys from Consul: %w", err)
	}

	if len(keys) == 0 {
		return nil, fmt.Errorf("no keys found with prefix '%s'", constant.PREFIX_CONSUL)
	}

	setValueConfig := &model.ConsulConfigRes{}

	// Iterasi melalui daftar kunci
	for _, kvPair := range keys {
		key := kvPair.Key
		value := kvPair.Value
		log.Printf("Processing key: %s", key)

		switch {
		case key == fmt.Sprintf("%s/appName", constant.PREFIX_CONSUL):
			setValueConfig.AppName = string(value)
		case key == fmt.Sprintf("%s/serverPort", constant.PREFIX_CONSUL):
			setValueConfig.ServerPort = string(value)
		case key == fmt.Sprintf("%s/environment", constant.PREFIX_CONSUL):
			setValueConfig.Environment = string(value)
		}

		parts := strings.Split(key, "/")
		if len(parts) < 3 {
			continue
		}

		envType := parts[len(parts)-2]
		if env == envType {
			configType := parts[len(parts)-3]
			contentType := parts[len(parts)-1]

			// Mengidentifikasi bagian dari key dan mengupdate map yang sesuai
			switch {
			case configType == "logDirectory":
				if contentType == "path" {
					setValueConfig.LogDirectory = string(value)
				}
			case configType == "redis":
				if contentType == "host" {
					setValueConfig.RedisHost = string(value)
				} else if contentType == "port" {
					setValueConfig.RedisPort = string(value)
				} else if contentType == "maxSize" {
					setValueConfig.RedisMaxSize = string(value)
				} else if contentType == "minIdleSize" {
					setValueConfig.RedisMinIdleSize = string(value)
				}
			case configType == "database":
				if contentType == "connection" {
					setValueConfig.DbConnection = string(value)
				} else if contentType == "maxPollLifeTime" {
					setValueConfig.DbMaxPollLifeTime = string(value)
				} else if contentType == "maxPoolIdle" {
					setValueConfig.DbMaxPoolIdle = string(value)
				} else if contentType == "maxPoolOpen" {
					setValueConfig.DbMaxPoolOpen = string(value)
				} else if contentType == "password" {
					setValueConfig.DbPassword = string(value)
				} else if contentType == "port" {
					setValueConfig.DbPort = string(value)
				} else if contentType == "schema" {
					setValueConfig.DbSchema = string(value)
				} else if contentType == "url" {
					setValueConfig.DbUrl = string(value)
				} else if contentType == "username" {
					setValueConfig.DbUsername = string(value)
				}
			case configType == "key":
				if contentType == "apiKey" {
					setValueConfig.ApiKey = string(value)
				} else if contentType == "apiKeyEncode" {
					setValueConfig.ApiKeyEncode = string(value)
				} else if contentType == "signatureKey" {
					setValueConfig.SignatureKey = string(value)
				} else if contentType == "signatureKeyEncode" {
					setValueConfig.SignatureKeyEncode = string(value)
				}
			case configType == "route":
				if contentType == "name" {
					setValueConfig.Route = string(value)
				}
			}
		}

	}

	return setValueConfig, nil
}
