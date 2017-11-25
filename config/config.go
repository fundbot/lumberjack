package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config object
type Config struct {
	name     string
	port     int
	logLevel string
	baseURL  string
}

var config *Config

// Load config from file
func Load() {
	viper.SetDefault("APP_PORT", "3000")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file %s was edited, reloading config\n", e.Name)
		readLatestConfig()
	})

	readLatestConfig()

}

func readLatestConfig() {
	config = &Config{
		name:     viper.GetString("application.name"),
		port:     viper.GetInt("application.port"),
		logLevel: viper.GetString("application.logLevel"),
		baseURL:  viper.GetString("reading.baseURL"),
	}

	fmt.Println(config)
}

// Application : Exporting configuration
func Application() *Config {
	return config
}

// Port : Exporting port
func Port() int {
	return config.port
}

// LogLevel : Exporting logLevel
func LogLevel() string {
	return config.logLevel
}

// BaseURL : Exporting baseURL
func BaseURL() string {
	return config.baseURL
}
