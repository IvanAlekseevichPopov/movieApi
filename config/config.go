package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	HTTPS bool   `default:"false" env:"HTTPS"`
	Port  string `default:":9000" env:"PORT"`
	DB    struct {
		Name     string `env:"DBName" default:"qor_example"`
		Adapter  string `env:"DBAdapter" default:"mysql"`
		Host     string `env:"DBHost" default:"localhost"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser"`
		Password string `env:"DBPassword"`
	}
}{}

func init() {
	if err := configor.Load(&Config, "config/database.yml", "config/application.yml"); err != nil {
		panic(err)
	}
}
