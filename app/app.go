package app

import (
	"net/http"
)

type Config struct {
	Addr string
}

type Server struct {
	cfg Config
	mux *http.ServeMux
}

func NewServer(cfg Config) (*Server, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("openmock"))
	})

	return &Server{
		cfg: cfg,
		mux: mux,
	}, nil
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.cfg.Addr, s.mux)
}
