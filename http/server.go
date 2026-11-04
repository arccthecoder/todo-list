package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	HTTPHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		HTTPHandlers: httpHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.HTTPHandlers.HandleCreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.HTTPHandlers.HandleGetTask)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.HTTPHandlers.HandleGetAllTasks)
	router.Path("/tasks").Methods("GET").Queries("completed", "true").HandlerFunc(s.HTTPHandlers.HandleGetAllUncompletedTasks)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.HTTPHandlers.HandleCompleteTask)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.HTTPHandlers.HandleDeleteTask)

	return http.ListenAndServe(":9091", router)
}
