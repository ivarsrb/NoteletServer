package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ivarsrb/NoteletServer/notes"
)

func newRouter() *gin.Engine {
	router := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	// Serve frontend static files
	//router.Use(static.Serve("/", static.LocalFile("./web", true)))
	// This may have problems when we need js files served, look in github help for other examples
	router.StaticFile("/", "./web/index.html")

	// Setup route group for the API
	api := router.Group("/api")
	{
		// Authorization middleware
		// api.Use(AuthRequired())
		api.GET("/notes", notes.GetNotes)
		api.GET("/notes/:id", notes.GetNote)
		api.POST("/notes", notes.PostNote)
		api.DELETE("/notes/:id", notes.DeleteNote)
	}

	return router
}
