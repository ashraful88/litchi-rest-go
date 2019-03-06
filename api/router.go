package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MountRoute func will mount all rest routes
func MountRoute(router *gin.RouterGroup) {
	router.GET("/health", healthCheck)
	router.POST("/entity", createEntity)
	router.GET("/entity", readEntityList)
	router.OPTIONS("/*any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "pass"})
	})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "pass"})
}

func createEntity(c *gin.Context) {
	c.JSON(handleCreate(c))
}

func readEntityList(c *gin.Context) {
	c.JSON(handleRead(c))
}
