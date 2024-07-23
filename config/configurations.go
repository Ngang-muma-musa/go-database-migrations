package config

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	MysqlConnection string `mapstructure:"MYSQL_CONNECTION" validate:"required"`
	Port            uint   `mapstructure:"PORT" validate:"required"`
}

func NewConfiguration() (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(wd + "/.env")
	if err == nil {
		err = godotenv.Load()
	}
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	var config Config
	if err := bindAll(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func bindAll(i interface{}) error {
	r := reflect.TypeOf(i)
	for r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	for i := 0; i < r.NumField(); i++ {
		env := r.Field(i).Tag.Get("mapstructure")
		err := viper.BindEnv(env)
		if err != nil {
			return err
		}
	}
	return viper.Unmarshal(i)
}
