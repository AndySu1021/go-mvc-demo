package main

import (
	"github.com/spf13/viper"
	syslog "log"
	"mvc/databases"
	"mvc/models"
	"mvc/routes"
	"mvc/utils/log"
	"strconv"
)

func main() {
	var err error

	// Init viper config
	err = initViper()
	if err != nil {
		return
	}

	// Init logger
	err = log.InitFluentd()
	if err != nil {
		syslog.Println("Init fluentd failed.")
	}
	defer log.FluentdClient.Close()

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
	databases.InitRedis()

	// Register the routes
	r := routes.SetRouter()

	// start server on 8080 port
	err = r.Run()
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

func reverse(x int) int {
	var result []rune
	if x < 0 {
		x *= -1
		result = append(result, '-')
	}
	tmp := strconv.Itoa(x)
	check := false
	for i:= len(tmp)-1; i >=0; i-- {
		if !check && tmp[i] == '0' {
			continue
		} else {
			check = true
			result = append(result, rune(tmp[i]))
		}
	}

	answer, _ := strconv.Atoi(string(result))
	return answer
}

