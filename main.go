package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"github.com/qor/qor"
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
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&User{}, &Staff{})
	DB.LogMode(true) //TOOD remove for prod

	go public(DB)

	// Initalize
	//Admin := admin.New(&admin.AdminConfig{DB: DB})
	Admin := admin.New(&qor.Config{DB: DB})

	// Allow to use Admin to manage User, Product
	Admin.AddResource(&User{})
	Admin.AddResource(&Staff{})

	// initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// Mount admin interface to mux
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)
}

func public(DB *gorm.DB) {
	r := gin.Default()

	r.GET("/staff", func(c *gin.Context) {
		fmt.Println(c.Request.URL.Query())
		name := c.Query("name")
		fmt.Println(name)

		var staffCollection []Staff
		DB.Where("name LIKE ?", "%"+name+"%").Find(&staffCollection)
		fmt.Println(staffCollection)

		c.JSON(http.StatusOK, staffCollection)
	})

	r.Run(":8080")
}
