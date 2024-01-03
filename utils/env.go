package utils

import (
	"github.com/spf13/viper"
	"path/filepath"
	"reflect"
	"runtime"
)

// Env has environment stored
type Env struct {
	ServerPort    string `mapstructure:"SERVER_PORT"`
	OauthClientID string `mapstructure:"OAUTH_CLIENT_ID"`
	OauthEndpoint string `mapstructure:"OAUTH_ENDPOINT"`
	DbHost        string `mapstructure:"DB_HOST"`
	DbUser        string `mapstructure:"DB_USER"`
	DbPass        string `mapstructure:"DB_PASS"`
	DbName        string `mapstructure:"DB_NAME"`
	DbPort        int    `mapstructure:"DB_PORT"`
}

// NewEnv creates a new environment
func NewEnv(log Logger) Env {

	env := Env{}

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	viper.SetConfigFile(filepath.Join(dir, "..", ".env"))
	viper.SetConfigType("env")

	// Read the config file (.env), but ignore the error if the file does not exist
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// The error is something other than the file not found
			log.Errorf("☠️ cannot read configuration: %v", err)
		}
	}

	v := reflect.ValueOf(env)
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		key := field.Tag.Get("mapstructure")
		if key == "" {
			key = field.Name
		}
		viper.BindEnv(key)
	}

	// Unmarshal the environment variables into the Env struct
	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	log.Infof("%+v \n", env)
	return env
}
