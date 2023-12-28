package utils

import (
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

// Env has environment stored
type Env struct {
	ServerPort    string `mapstructure:"SERVER_PORT"`
	OauthClientID string `mapstructure:"OAUTH_CLIENT_ID"`
	OauthEndpoint string `mapstructure:"OAUTH_ENDPOINT"`
}

// NewEnv creates a new environment
func NewEnv(log Logger) Env {

	env := Env{}
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	viper.SetConfigFile(filepath.Join(dir, "..", ".env"))
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	log.Infof("%+v \n", env)
	return env
}
