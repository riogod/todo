package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
	apiv1 "todo_api/internal/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	model "github.com/riogod/todo/libs/gomodels"
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
	Router     *gin.Engine
	httpServer *http.Server
}

func NewAppInit() *App {
	config, err := readConfig()

	if err != nil {
		panic("Cannot loading configs from env")
	}
	db := initDB(*config.DatabaseConfig)

	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	apiv1.Setup(router, db)

	return &App{
		Config: config,
		DB:     db,
		Router: router,
	}
}

func (app *App) Run(router *gin.Engine) error {

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
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
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

	db.AutoMigrate(&model.ToDoItemList{})

	return db
}
