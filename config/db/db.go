package db

import (
	"errors"
	"fmt"
	"github.com/IvanAlekseevichPopov/movieApi/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/l10n"
	"github.com/qor/media"
	"github.com/qor/publish2"
	"github.com/qor/sorting"
	"github.com/qor/validations"
	"os"
)

// Conn Global Conn connection
var Conn *gorm.DB

func init() {
	var err error

	dbConfig := config.Config.DB
	if config.Config.DB.Adapter == "mysql" {
		Conn, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
		// Conn = Conn.Set("gorm:table_options", "CHARSET=utf8")
	} else if config.Config.DB.Adapter == "postgres" {
		Conn, err = gorm.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name))
	} else if config.Config.DB.Adapter == "sqlite" {
		Conn, err = gorm.Open("sqlite3", fmt.Sprintf("./%v", dbConfig.Name))
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		if os.Getenv("DEBUG") != "" {
			Conn.LogMode(true)
		}

		l10n.RegisterCallbacks(Conn)
		sorting.RegisterCallbacks(Conn)
		validations.RegisterCallbacks(Conn)
		media.RegisterCallbacks(Conn)
		publish2.RegisterCallbacks(Conn)
	} else {
		panic(err)
	}
}
