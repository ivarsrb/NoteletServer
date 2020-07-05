package server

import (
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	// staticPath points to directory of static files to be served
	staticPath = "./web"
	// REST api urls
	apiURL   = "/api"
	notesURL = "/notes"
)

// newRouter returns a router for REST api and SPA client
func newRouter() *gin.Engine {
	router := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	// Limit the maximal size of request srecieved, error is logged otherwie
	const maxBodySize = 1 << 20
	router.Use(limits.RequestSizeLimiter(maxBodySize))
	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile(staticPath, true)))
	// Setup route group for the REST API
	api := router.Group(apiURL)
	{
		// Authorization middleware
		// api.Use(AuthRequired())
		api.GET(notesURL, getNotes)
		api.GET(notesURL+"/:id", getNote)
		api.POST(notesURL, postNote)
		api.DELETE(notesURL+"/:id", deleteNote)
		api.PUT(notesURL+"/:id", replaceNote)
	}

	return router
}
