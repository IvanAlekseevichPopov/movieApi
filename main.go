package main

import (
	"fmt"
	"github.com/IvanAlekseevichPopov/movieApi/config"
	"github.com/IvanAlekseevichPopov/movieApi/config/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"log"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name string
}

type Staff struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"-"`
	DeletedAt      *time.Time `sql:"index" json:"-"`
	Name           string     `json:"name"`
	Birthday       time.Time  `json:"birthday"`
	IsActor        bool       `json:"isActor"`
	IsProducer     bool       `json:"isProducer"`
	IsDirector     bool       `json:"isDirector"`
	IsScreenWriter bool       `json:"isScreenWriter"`
}

func main() {
	db.Conn.AutoMigrate(User{}, Staff{}) //TODO only separate migrations

	// initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	api(mux)
	adminPanel(mux)

	log.Printf("Listening on%s", config.Config.Port)
	http.ListenAndServe(config.Config.Port, mux)
}

func api(mux *http.ServeMux) {
	router := gin.Default()

	router.GET("/api/staff", func(c *gin.Context) {
		fmt.Println(c.Request.URL.Query())
		name := c.Query("name")
		fmt.Println(name)

		var staffCollection []Staff
		db.Conn.Where("name LIKE ?", "%"+name+"%").Find(&staffCollection)
		fmt.Println(staffCollection)

		if len(staffCollection) == 0 {
			c.JSON(http.StatusNotFound, gin.H{}) //TODO формат ответов
			//TODO проверка черных списков
			//TODO Запуск поиска по ресурсам
		} else {
			c.JSON(http.StatusOK, staffCollection)
		}

	})

	mux.Handle("/api/staff", router) //TODO забиндить все маршруты /api*
}

func adminPanel(mux *http.ServeMux) {
	// Initalize
	Admin := admin.New(&admin.AdminConfig{DB: db.Conn})

	// Allow to use Admin to manage User, Product
	Admin.AddResource(&User{})
	Admin.AddResource(&Staff{})

	// Mount admin interface to mux
	Admin.MountTo("/admin", mux)
}
