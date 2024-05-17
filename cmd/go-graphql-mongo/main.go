package main

import (
	"net/http"
	"time"

	"github.com/pi-prakhar/go-graphql-mongo/api"
	"github.com/pi-prakhar/go-graphql-mongo/internal/config/db"
	"github.com/pi-prakhar/go-graphql-mongo/pkg/logger"
	"github.com/pi-prakhar/go-graphql-mongo/pkg/utils"
	"github.com/pi-prakhar/utils/loader"
	"github.com/rs/cors"
)

func init() {
	logger.InitLogger()
	logger.Log.Info("GO-PHONE-OTP-SERVICE Logger Started")

	err := loader.LoadEnv()
	if err != nil {
		logger.Log.Error("Failed to Load ENV", err)
	}

	db.Connect()
}

func main() {
	hostAddress := utils.GetHostAddress()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"localhost:3000", "pi-go-backend:8000", "localhost:8000"},
		AllowedMethods:   []string{"POST", "GET"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler(api.Router())

	srv := &http.Server{
		Handler:      corsHandler,
		Addr:         hostAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Log.Info("connect to localhost:8000/playground for GraphQL playground")
	logger.Log.Error("Error : Failed to start server", srv.ListenAndServe())
}
