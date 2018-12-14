package controller

import (
	"fmt"
	"github.com/IvanAlekseevichPopov/movieApi/config/db"
	"github.com/IvanAlekseevichPopov/movieApi/entity"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"net/http"
)

func NewApi(router *mux.Router, route string) {
	ginRouter := gin.Default()

	//TODO убрать известный префикс /api
	ginRouter.GET("/api/staff", func(c *gin.Context) {
		name := c.Query("name")

		var staffCollection []entity.Staff
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

	router.PathPrefix(route).Handler(ginRouter)
}
