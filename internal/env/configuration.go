package env

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

var (
	once   sync.Once
	config = &configuration{}
)

type configuration struct {
	App                 App                 `json:"app"`
	DB                  DB                  `json:"db"`
	AuthService         AuthService         `json:"auth_service"`
	TransactionsService TransactionsService `json:"transactions_service"`
	BlockService        BlockService        `json:"block_service"`
}

type App struct {
	ServiceName       string `json:"service_name"`
	PathLog           string `json:"path_log"`
	LogReviewInterval int    `json:"log_review_interval"`
	Language          string `json:"language"`
	UserLogin         string `json:"user_login"`
	UserPassword      string `json:"user_password"`
	TimerInterval     string `json:"timer_interval"`
	SubscriptionTime  int    `json:"subscription_time"`
	MaxMiners         int    `json:"max_miners"`
	MaxValidator      int    `json:"max_validator"`
	FeeMine           int    `json:"fee_mine"`
	FeeValidators     int    `json:"fee_validators"`
	WalletMain        string `json:"wallet_main"`
}

type DB struct {
	Engine   string `json:"engine"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Instance string `json:"instance"`
	IsSecure bool   `json:"is_secure"`
	SSLMode  string `json:"ssl_mode"`
}

type AuthService struct {
	Port string `json:"port"`
}

type TransactionsService struct {
	Port string `json:"port"`
}

type BlockService struct {
	Port string `json:"port"`
}

func NewConfiguration() *configuration {
	fromFile()
	return config
}

// LoadConfiguration lee el archivo configuration.json
// y lo carga en un objeto de la estructura Configuration
func fromFile() {
	once.Do(func() {
		b, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Fatalf("no se pudo leer el archivo de configuración: %s", err.Error())
		}

		err = json.Unmarshal(b, config)
		if err != nil {
			log.Fatalf("no se pudo parsear el archivo de configuración: %s", err.Error())
		}

		if config.DB.Engine == "" {
			log.Fatal("no se ha cargado la información de configuración")
		}
	})
}
