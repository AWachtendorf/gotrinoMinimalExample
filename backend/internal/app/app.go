package app

import (
	"rochus/backend/internal/httpserver"
)

type Application struct {
	httpPort          int
	httpServer        *httpserver.HttpServer

}

// getHttpServer returns a instance of httpserver.HttpServer. When the
// instance in Application is nil, a new one will created.
func (a *Application) getHttpServer() *httpserver.HttpServer {
	if a.httpServer == nil {
		a.httpServer = httpserver.NewHttpServer(a.httpPort)
	}

	return a.httpServer
}






// Run is used, to start the usecases
func (a *Application) Run() {


	a.httpServer.Start()
}

// NewApplication creates a new instance of Application and returns it
func NewApplication(httpPort int) *Application {
	return &Application{
		httpPort: httpPort,
	}
}
