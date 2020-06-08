package server

import (
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

// newRouter describes a router for REST api and SPA client
func newRouter() *gin.Engine {
	router := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	// Limit the maximal size of request srecieved, error is logged otherwie
	const maxBodySize = 1048576
	router.Use(limits.RequestSizeLimiter(maxBodySize))
	// Serve frontend static files
	//router.Use(static.Serve("/", static.LocalFile("./web", true)))
	// TODO: This may have problems when we need js files served, look in github help for other examples
	router.StaticFile("/", "./web/index.html")

	// Setup route group for the REST API
	api := router.Group("/api")
	{
		// Authorization middleware
		// api.Use(AuthRequired())
		api.GET("/notes", getNotes)
		api.GET("/notes/:id", getNote)
		api.POST("/notes", postNote)
		api.DELETE("/notes/:id", deleteNote)
	}

	return router
}
