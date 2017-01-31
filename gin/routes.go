package gin

import "github.com/gin-gonic/gin"

// SetRoutes : sets routes for all gin
func SetRoutes(router *gin.Engine) {
	router.GET("/", indexHandler)
}
