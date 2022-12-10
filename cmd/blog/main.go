package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/dhevve/blog"
	"github.com/dhevve/blog/internal/handler"
	"github.com/dhevve/blog/internal/repository"
	"github.com/dhevve/blog/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var validate *validator.Validate
var ctx = context.Background()

func main() {
	if err := intiConfig(); err != nil {
		logrus.Fatalf("error config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("dotenv %s", err.Error())
	}

	validate = validator.New()
	uri := "neo4j://localhost:7687"
	auth := neo4j.BasicAuth("neo4j", "root", "")
	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("fail: %s", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.DB"),
	})

	repo := repository.NewRepository(db, driver)
	services := service.NewService(repo, redisClient)
	handlers := handler.NewHandler(services, validate, ctx)
	srv := new(blog.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func intiConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func unsafeClose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		logrus.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}
