package app

import (
	"net/http"
)

type Server struct {
	Addr string
	mux  *http.ServeMux
}

func NewServer() (*Server, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("openmock"))
	})

	return &Server{
		Addr: ":8080",
		mux:  mux,
	}, nil
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.Addr, s.mux)
}
