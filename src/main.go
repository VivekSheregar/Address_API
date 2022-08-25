package main

import (
	"context"

	"github.com/VivekSheregar/address_API/utils"
	"github.com/VivekSheregar/address_API/views"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := utils.GetConnection(context.Background(), "postgres", "1980", "localhost", "5432")
	router := gin.Default()
	s := views.NewStateView(conn)
	router.POST("/states", s.Create)
	router.GET("/states", s.List)
	router.GET("/states/:id", s.Get)
	router.PUT("/states/:id", s.Update)
	router.DELETE("/states/:id", s.Delete)
	router.DELETE("/states/all", s.DeleteAll)
	router.Run()
}
