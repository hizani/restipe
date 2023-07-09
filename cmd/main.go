package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"restipe/internal/handler"
	"restipe/internal/server"
	"restipe/internal/service"
	"restipe/internal/storage"
	"restipe/internal/storage/sqldb"
	"syscall"

	"github.com/spf13/viper"
)

// @title Recipe App API
// @version 0.1
// @description API Server for Recipe Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	configDirPath := os.Args[1]

	viper.AddConfigPath(configDirPath)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {

		log.Fatalf("can't initialize config: %s", err.Error())
	}

	storage, err := storage.NewPostgres(sqldb.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.ssl"),
	})
	if err != nil {
		log.Fatalf("can't initialize storage: %s", err.Error())
	}
	service := service.New(storage, viper.GetString("jwttoken"))
	handler := handler.NewGinHandler(service)

	server := server.New(viper.GetString("port"), handler)
	go func() {
		if err := server.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("unexpected server shutdown: %s", err.Error())
		}
	}()

	sigQuit := make(chan os.Signal)
	signal.Notify(sigQuit, syscall.SIGTERM, syscall.SIGINT)
	<-sigQuit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("error on server shutdown: %s", err.Error())
	}

	if err := storage.Close(); err != nil {
		log.Printf("error on db close: %s", err.Error())
	}

}
