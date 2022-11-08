package api

import (
	"fmt"
	"net/http"
	"time"

	"npm/internal/logger"
)

const httpPort = 3000

// StartServer creates a http server
func StartServer() {
	logger.Info("Server starting on port %v", httpPort)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%v", httpPort),
		Handler:           NewRouter(),
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error("HttpListenError", err)
	}
}