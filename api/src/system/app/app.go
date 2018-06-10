package app

import (
	"log"
	"net/http"
)

// just a struct for the seever
type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}

// initialize server
func (S *Server) Init(port string) {
	log.Println("Initializing server...")
	log.Println(port)
	S.port = ":" + port
}

//start server
func (S *Server) Start() {
	log.Println("starting server on port" + S.port)
	http.ListenAndServe(S.port, nil)
}
