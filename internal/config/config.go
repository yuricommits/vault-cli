package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	KeyToken   = "token"
	KeyBaseURL = "base_url"
)

func Init() {
	configDir, _ := os.UserConfigDir()
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(filepath.Join(configDir, "vault"))
	viper.SetDefault(KeyBaseURL, "https://vault-two-lovat.vercel.app")
	viper.ReadInConfig()
}

func Save() error {
	configDir, _ := os.UserConfigDir()
	dir := filepath.Join(configDir, "vault")
	os.MkdirAll(dir, 0700)
	viper.SetConfigFile(filepath.Join(dir, "config.json"))
	return viper.WriteConfig()
}

func GetToken() string {
	return viper.GetString(KeyToken)
}

func GetBaseURL() string {
	return viper.GetString(KeyBaseURL)
}

func SetToken(token string) {
	viper.Set(KeyToken, token)
}

func SetBaseURL(url string) {
	viper.Set(KeyBaseURL, url)
}

func IsAuthenticated() bool {
	return GetToken() != ""
}
