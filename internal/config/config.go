package config

import (
	ui "github.com/YashKumarVerma/go-lib-ui"
	"github.com/spf13/viper"
)

// load configurations
func loadConfigurations() {
	viper.SetConfigName("default")
	viper.AddConfigPath("./config/")
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	ui.CheckError(err, "config file not found, using default config", false)

	// set default configurations
	viper.SetDefault("app.name", "Reverse Coding : Code Runner")
	viper.SetDefault("app.code_directory", "codes")
	viper.SetDefault("app.port", 8000)
	viper.SetDefault("app.seed_source", "https://rc-broadcaster.vercel.app/api/app.js")

	// save configurations onto exported object
	Load.Name = viper.GetString("app.name")
	Load.CodeDirectory = viper.GetString("app.code_directory")
	Load.Port = viper.GetInt("app.port")
	Load.SeedSource = viper.GetString("app.seed_source")
}

// Configuration store
type Configuration struct {
	Name          string
	CodeDirectory string
	SeedSource    string
	Port          int
}

// Load configurations to be used from other modules
var Load Configuration

// Initialize config parser
func Initialize() {
	ui.ContextPrint("key", "Loading configurations")
	loadConfigurations()
}
