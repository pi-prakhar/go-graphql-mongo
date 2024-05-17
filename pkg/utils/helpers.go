package utils

import (
	"github.com/pi-prakhar/go-graphql-mongo/pkg/logger"
	"github.com/pi-prakhar/utils/loader"
)

func GetHostAddress() string {
	var hostAddress string
	isDocker, err := loader.GetValueFromConf("docker")
	if err != nil {
		logger.Log.Error("Error : Failed to fetch docker property from configurations", err)
	}

	if isDocker == "true" {
		hostAddress, err = loader.GetValueFromConf("docker-host-address")
		if err != nil {
			logger.Log.Error("Error : Failed to fetch docker-host-address property from configurations", err)
		}
	} else {
		hostAddress, err = loader.GetValueFromConf("local-host-address")
		if err != nil {
			logger.Log.Error("Error : Failed to fetch local-host-address property from configurations", err)
		}
	}
	return hostAddress
}
