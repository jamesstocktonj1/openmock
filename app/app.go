package app

import (
	"net/http"
)

type Server struct {
	Addr string
}

func (s *Server) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("openmock"))
	})

	return http.ListenAndServe(s.Addr, mux)
}
