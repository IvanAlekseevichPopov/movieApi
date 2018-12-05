//package main
//
//import (
//	"fmt"
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/sqlite"
//)
//
//
//type Person struct {
//	ID uint
//	FirstName string
//	LastName string
//}
//
//var db *gorm.DB
//
//func setupRouter() *gin.Engine {
//	r := gin.Default()
//
//	// Ping test
//	r.GET("/ping", func(c *gin.Context) {
//		c.String(http.StatusNoContent, "pong")
//	})
//
//	r.GET("/tst", func(c *gin.Context) {
//		var user Person
//		db.First(user)
//		fmt.Println(db)
//		c.JSON(http.StatusOK, gin.H{"user": "user"})
//	})
//
//
//	return r
//}
//
//func main() {
//	db, _ = gorm.Open("sqlite3", "./gorm.db")
//	db.AutoMigrate(&Person{})
//	p1 := Person{FirstName: "John", LastName: "Doe"}
//	db.Create(&p1)
//	defer db.Close()
//
//	r := setupRouter()
//	// Listen and Server in 0.0.0.0:8080
//	r.Run(":8080")
//}

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

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusNoContent, "pong")
	})

	r.GET("/staff", func(c *gin.Context) {
		name := c.Query("name")
		fmt.Println(name)

		var staffCollection []Staff
		DB.Where("name LIKE ?", name).Find(&staffCollection) //TODO like with %, parameter validation
		fmt.Println(staffCollection)

		c.JSON(http.StatusOK, staffCollection)
	})

	r.Run(":8080")
}
