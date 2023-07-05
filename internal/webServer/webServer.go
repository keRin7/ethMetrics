package webServer

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/sirupsen/logrus"
)

type WebServer struct {
	config *Config
}

func CreateWebServer(config *Config) *WebServer {
	return &WebServer{
		config: config,
	}
}

func (r *WebServer) AddHeandler(pattern string, handler http.Handler) {
	r.config.mux.Handle(pattern, handler)
}

func (r *WebServer) Start() {
	server := http.Server{
		Addr:         ":" + r.config.Port,
		Handler:      r.config.mux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	logrus.Info("Service starting on port:", r.config.Port, "...")
	if err := server.ListenAndServe(); err != nil {
		logrus.Println(err)
	}
}
