package main

import (
	"github.com/VivekSheregar/address_API/views"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	s := views.NewStateView()
	router.POST("/states", s.Create)
	router.GET("/states", s.List)
	router.GET("/states/:id", s.Get)
	router.PUT("/states/:id", s.Update)
	router.DELETE("/states/:id", s.Delete)
	router.Run()
}
