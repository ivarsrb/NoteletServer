package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

// Create creates and returns handler object
// that is concerned with routing of server request paths
func createRouter() *mux.Router {
	/*
		// Router
		router := mux.NewRouter()
		// REST api serving
		apirouter := router.PathPrefix("/api/notes").Subrouter()
		apirouter.HandleFunc("", getNotes).Methods(http.MethodGet)
		apirouter.HandleFunc("/{id:[0-9]+}", getNote).Methods(http.MethodGet)
		apirouter.HandleFunc("", postNote).Methods(http.MethodPost)
		apirouter.HandleFunc("/{id:[0-9]+}", deleteNote).Methods(http.MethodDelete)
		// SPA serving
		spa := spaHandler{staticPath: "static", indexPath: "index.html"}
		router.PathPrefix("/").Handler(spa)
		return router
	*/
	return nil
}

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
		api.GET("/notes", getNotes)
		api.GET("/notes/:id", getNote)
		api.POST("/notes", postNote)
		api.DELETE("/notes/:id", deleteNote)
	}

	return router
}
