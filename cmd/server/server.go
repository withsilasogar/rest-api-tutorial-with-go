package server

import (
	"context"
	"net/http"
	"test01/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GinServer interface defines methods for starting, shutting down, and registering routes in the Gin server.
type GinServer interface {
	Start(ctx context.Context, httpAddress string) error
	Shutdown(ctx context.Context) error
	RegisterGroupRoute(path string, routes []interfacesx.RouteDefinition, middleWare ...gin.HandlerFunc)
	RegisterRoute(method, path string, handler gin.HandlerFunc)
}

// GinServerBuilder is an empty struct used for building a Gin server.
type GinServerBuilder struct {
}

// ginServer represents the Gin server implementation.
type ginServer struct {
	engine *gin.Engine  // Gin engine instance
	server *http.Server // HTTP server instance
}

// NewGinServerBuilder creates and returns a new instance of GinServerBuilder.
func NewGinServerBuilder() *GinServerBuilder {
	return &GinServerBuilder{}
}

// Build creates a new ginServer instance with a default Gin engine.
func (b *GinServerBuilder) Build() GinServer {
	engine := gin.Default()
	return &ginServer{engine: engine}
}

// Start initializes and starts the HTTP server.
func (gs *ginServer) Start(ctx context.Context, httpAddress string) error {
	gs.server = &http.Server{
		Addr:    httpAddress, // Server address
		Handler: gs.engine,   // Gin engine as the HTTP handler
	}

	// Start the server in a new goroutine
	go func() {
		if err := gs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Listening %s \n", err) // Log fatal error if server fails to start
		}
	}()

	logrus.Infof("Https server is running on Port %s", httpAddress) // Log info indicating server is running
	return nil
}

// Shutdown gracefully shuts down the HTTP server.
func (gs *ginServer) Shutdown(ctx context.Context) error {
	if err := gs.server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server shutdown %s", err) // Log fatal error if shutdown fails
	}

	logrus.Info("Server is exiting") // Log info indicating server is exiting

	return nil
}

// RegisterRoute registers a single route with the specified method, path, and handler.
func (gs *ginServer) RegisterRoute(method, path string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		gs.engine.GET(path, handler)
	case "POST":
		gs.engine.POST(path, handler)
	case "PUT":
		gs.engine.PUT(path, handler)
	case "DELETE":
		gs.engine.DELETE(path, handler)
	case "PATCH":
		gs.engine.PATCH(path, handler)
	default:
		logrus.Errorf("Invalid https method") // Log error if method is invalid
	}
}

// RegisterGroupRoute registers multiple routes under a group with the specified path and middleware.
func (gs *ginServer) RegisterGroupRoute(path string, routes []interfacesx.RouteDefinition, middleWare ...gin.HandlerFunc) {
	group := gs.engine.Group(path) // Create a route group with the specified path
	group.Use(middleWare...)       // Apply middleware to the group
	for _, route := range routes { // Iterate over the provided route definitions
		switch route.Method {
		case "GET":
			group.GET(route.Path, route.Handler)
		case "POST":
			group.POST(route.Path, route.Handler)
		case "PUT":
			group.PUT(route.Path, route.Handler)
		case "DELETE":
			group.DELETE(route.Path, route.Handler)
		case "PATCH":
			group.PATCH(route.Path, route.Handler)
		default:
			logrus.Errorf("Invalid https method") // Log error if method is invalid
		}
	}
}
