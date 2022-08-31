package server

import (
	// import third-party packages
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// import local packages
	"github.com/impossible98/grayfox/internal/server/api"
)

func InitServer() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/version", api.Version)
	router.Run()
}
