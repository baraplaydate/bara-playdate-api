package model

type LogDirectoryReq struct {
	Path string `yaml:"path"`
}

type RedisConfigReq struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	MaxSize     string `yaml:"maxSize"`
	MinIdleSize string `yaml:"minIdleSize"`
}

type DatabaseConfigReq struct {
	Connection      string `yaml:"connection"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	URL             string `yaml:"url"`
	Port            string `yaml:"port"`
	Schema          string `yaml:"schema"`
	MaxPoolOpen     string `yaml:"maxPoolOpen"`
	MaxPoolIdle     string `yaml:"maxPoolIdle"`
	MaxPollLifeTime string `yaml:"maxPollLifeTime"`
}

type KeyAccessReq struct {
	APIKey             string `yaml:"apiKey"`
	APIKeyEncode       string `yaml:"apiKeyEncode"`
	SignatureKey       string `yaml:"signatureKey"`
	SignatureKeyEncode string `yaml:"signatureKeyEncode"`
}

type RouteReq struct {
	Name string `yaml:"name"`
}

type ConsulConfigReq struct {
	AppName      string                       `yaml:"appName"`
	ServerPort   string                       `yaml:"serverPort"`
	Environment  string                       `yaml:"environment"`
	ReleaseMode  string                       `yaml:"releaseMode"`
	LogDirectory map[string]LogDirectoryReq   `yaml:"logDirectory"`
	Redis        map[string]RedisConfigReq    `yaml:"redis"`
	Database     map[string]DatabaseConfigReq `yaml:"database"`
	Key          map[string]KeyAccessReq      `yaml:"key"`
	Route        map[string]RouteReq          `yaml:"route"`
}

type ConsulConfigRes struct {
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

// type ConsulConfig struct {
// 	AppName      string            `yaml:"appName"`
// 	ServerPort   string            `yaml:"serverPort"`
// 	Environment  string            `yaml:"environment"`
// 	ReleaseMode  string            `yaml:"releaseMode"`
// 	LogDirectory LogDirectory      `yaml:"logDirectory"`
// 	Redis        EnvironmentConfig `yaml:"redis"`
// 	Database     EnvironmentConfig `yaml:"database"`
// 	Key          EnvironmentKey    `yaml:"key"`
// 	Route        EnvironmentRoute  `yaml:"route"`
// }

// type LogDirectory struct {
// 	Dev  Directory `yaml:"dev"`
// 	QA   Directory `yaml:"qa"`
// 	UAT  Directory `yaml:"uat"`
// 	Prod Directory `yaml:"prod"`
// }

// type Directory struct {
// 	Path string `yaml:"path"`
// }

// type EnvironmentConfig struct {
// 	Dev  ServiceConfig `yaml:"dev"`
// 	QA   ServiceConfig `yaml:"qa"`
// 	UAT  ServiceConfig `yaml:"uat"`
// 	Prod ServiceConfig `yaml:"prod"`
// }

// type ServiceConfig struct {
// 	Host            string `yaml:"host"`
// 	Port            string `yaml:"port"`
// 	MaxSize         string `yaml:"maxSize"`
// 	MinIdleSize     string `yaml:"minIdleSize"`
// 	Connection      string `yaml:"connection"`
// 	Username        string `yaml:"username"`
// 	Password        string `yaml:"password"`
// 	URL             string `yaml:"url"`
// 	Schema          string `yaml:"schema"`
// 	MaxPoolOpen     string `yaml:"maxPoolOpen"`
// 	MaxPoolIdle     string `yaml:"maxPoolIdle"`
// 	MaxPollLifeTime string `yaml:"maxPollLifeTime"`
// }

// type EnvironmentKey struct {
// 	Dev  KeyConfig `yaml:"dev"`
// 	QA   KeyConfig `yaml:"qa"`
// 	UAT  KeyConfig `yaml:"uat"`
// 	Prod KeyConfig `yaml:"prod"`
// }

// type KeyConfig struct {
// 	APIKey             string `yaml:"apiKey"`
// 	APIKeyEncode       string `yaml:"apiKeyEncode"`
// 	SignatureKey       string `yaml:"signatureKey"`
// 	SignatureKeyEncode string `yaml:"signatureKeyEncode"`
// }

// type EnvironmentRoute struct {
// 	Dev  RouteConfig `yaml:"dev"`
// 	QA   RouteConfig `yaml:"qa"`
// 	UAT  RouteConfig `yaml:"uat"`
// 	Prod RouteConfig `yaml:"prod"`
// }

// type RouteConfig struct {
// 	Name string `yaml:"name"`
// }
