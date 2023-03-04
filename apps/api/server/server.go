package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string `required:"true"                                        envconfig:"DB_HOST"`
	Port     string `required:"true"                                        envconfig:"DB_PORT"`
	DBName   string `required:"true"                                        envconfig:"DB_NAME"`
	User     string `required:"true"                                        envconfig:"DB_USER"`
	Password string `required:"true"                                        envconfig:"DB_PASSWORD"`
	SSLMode  string `required:"true"                                        envconfig:"DB_SSL"`
}

type Config struct {
	DatabaseConfig *DatabaseConfig
	Debug          bool   `required:"true"  default:"false"                 envconfig:"DEBUG"`
	APP_ENV        string `required:"true" default:"development"            envconfig:"APP_ENV"`
	PORT           string `required:"true" default:"3030"            envconfig:"PORT"`
}

type App struct {
	Config     *Config
	DB         *gorm.DB
	httpServer *http.Server
}

func NewAppInit() *App {
	config, err := readConfig()

	if err != nil {
		panic("Cannot loading configs from env")
	}
	db := initDB(*config.DatabaseConfig)
	return &App{
		Config: config,
		DB:     db,
	}
}

func (app *App) Run() {

	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	app.httpServer = &http.Server{
		Addr:           ":" + app.Config.PORT,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()
}

func readConfig() (*Config, error) {
	var newCfg Config
	var err error

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	wd = filepath.Join(wd, "..", "..")

	envPath := filepath.Join(wd, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err = envconfig.Process("", &newCfg); err != nil {
		return nil, err
	}

	return &newCfg, nil
}

func initDB(dbConfig DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
		dbConfig.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error connect to database")
	}

	return db
}
