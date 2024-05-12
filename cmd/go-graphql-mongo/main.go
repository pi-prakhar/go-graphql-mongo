package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pi-prakhar/go-graphql-mongo/api"
	"github.com/pi-prakhar/go-graphql-mongo/internal/config/db"
	"github.com/pi-prakhar/go-graphql-mongo/pkg/logger"
	"github.com/pi-prakhar/utils/loader"
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
	hostAddress, err := loader.GetValueFromConf("local-host-address")
	if err != nil {
		logger.Log.Error("Error : Failed to load local-host-address from configuration", err)
	}

	srv := &http.Server{
		Handler:      api.Router(),
		Addr:         hostAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Log.Info(fmt.Sprintf("connect to http://localhost%s/playground for GraphQL playground", hostAddress))
	logger.Log.Error("Error : Failed to start server", srv.ListenAndServe())
}
