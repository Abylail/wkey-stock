package definition

import (
	"github.com/lowl11/lazyconfig/confapi"
	"github.com/lowl11/lazylog/logapi"
	"os"
)

type Configuration struct {
	Environment string

	Server struct {
		Port string `json:"port"`
	} `json:"server"`

	Database struct {
		Connection string `json:"connection"`
	} `json:"database"`

	Rabbit struct {
		Connection string `json:"connection"`
	} `json:"rabbit"`
}

var Config Configuration
var Logger logapi.ILogger

func Init() {
	Config = Configuration{}

	// определение окружения (прод или нет)
	Config.Environment = os.Getenv("env")
	isProduction := Config.Environment == "production"

	// создание логгера
	logger := logapi.New().File("info", "logs")

	// чтение конфигов
	if err := confapi.Read(&Config, isProduction); err != nil {
		logger.Fatal(err, "Reading config error")
	}

	Logger = logger

	// создание объекта сервера
	initServer()
}
