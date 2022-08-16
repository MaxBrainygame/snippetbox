package main

import "github.com/spf13/viper"

type config struct {
	port        string
	staticfiles string
}

func (app *application) newConfig() *config {

	viper.SetConfigName("config")
	viper.AddConfigPath("$Home/Projects/snippetbox")
	err := viper.ReadInConfig()
	if err != nil {
		app.errorLog.Fatal("fatal error config file: %w", err)
	}

	var config config
	config.port = viper.GetString("port")
	config.staticfiles = viper.GetString("staticfiles")

	return &config

}
