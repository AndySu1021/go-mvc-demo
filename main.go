package main

import (
	"github.com/spf13/viper"
	"mvc/databases"
	"mvc/models"
	"mvc/routes"
	"os"
)

func main() {
	var err error

	// Init viper config
	err = initViper()
	if err != nil {
		return
	}

	// Init logger
	//err = log.InitFluentd()
	//if err != nil {
	//	syslog.Println("Init fluentd failed.")
	//}
	//defer log.FluentdClient.Close()

	// Init database
	err = databases.InitMySql()
	if err != nil{
		return
	}
	defer databases.Close()

	// Make migration
	err = databases.MySqlClient.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	// Init Redis
	//databases.InitRedis()

	// Register the routes
	r := routes.SetRouter()

	// start server on 8080 port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = r.Run(":" + port)
	if err != nil {
		return
	}
}

func initViper() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	return
}
