package httpserver

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	cors2 "github.com/rs/cors"
	"log"
	"net/http"
	"strconv"
)

// HttpServer defines implements a backend component,
// to accept and handle the incoming http-requests with
// the given router
type HttpServer struct {
	router     *httprouter.Router
	port       int
}

// Router returns the router of the HttpServer instance
func (h *HttpServer) Router() *httprouter.Router {
	return h.router
}

// ReturnAsJson can be used, to return some data in the json format in the response
func (h *HttpServer) ReturnAsJson(data interface{}, writer http.ResponseWriter) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	writer.Header().Set("Content-Type", "usecases/json")
	if _, err := writer.Write([]byte(jsonData)); err != nil {
		return err
	}

	return err
}

// notFound will be called by the start, to handle the incoming
// requests, that are ending in a 404 error
func (h *HttpServer) notFound() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("The requested page can not be found."))
	})
}

// Start is used, to start the HttpServer on the given port
func (h *HttpServer) Start() {
	h.router.NotFound = h.notFound()



	cors := cors2.AllowAll().Handler(h.router)

	log.Println("Start server of port " + strconv.Itoa(h.port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(h.port), cors))
}

// NewHttpServer create a new instance of HttpServer
func NewHttpServer(port int) *HttpServer {
	return &HttpServer{
		router:     httprouter.New(),
		port:       port,
	}
}
