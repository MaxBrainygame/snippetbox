package main

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	port        string
	staticfiles string
	dsn         string
}

func newConfig(errorLog *log.Logger) *config {

	viper.SetConfigName("config")
	viper.AddConfigPath("$Home/Projects/snippetbox")
	err := viper.ReadInConfig()
	if err != nil {
		errorLog.Fatal("fatal error config file: %w", err)
	}

	var config config
	config.port = viper.GetString("port")
	config.staticfiles = viper.GetString("staticfiles")
	config.dsn = viper.GetString("dsn")

	return &config

}
