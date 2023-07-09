package main

import (
	"log"
	"os"
	"restipe/internal/handler"
	"restipe/internal/server"
	"restipe/internal/service"
	"restipe/internal/storage"
	"restipe/internal/storage/sqldb"

	"github.com/spf13/viper"
)

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
	if err := server.Run(); err != nil {
		log.Fatalf("unexpected server shutdown: %s", err.Error())
	}

}
