package config

import (
	"fmt"
	"npm/internal/logger"
	"os"
)

// createDataFolders will recursively create these folders within the
// data folder defined in configuration.
func createDataFolders() {
	folders := []string{
		"access",
		"certificates",
		"logs",
		// Acme.sh:
		Configuration.Acmesh.GetWellknown(),
		// Nginx:
		"nginx/hosts",
		"nginx/streams",
		"nginx/temp",
	}

	for _, folder := range folders {
		path := folder
		if path[0:1] != "/" {
			path = fmt.Sprintf("%s/%s", Configuration.DataFolder, folder)
		}
		logger.Debug("Creating folder: %s", path)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			logger.Error("CreateDataFolderError", err)
		}
	}
}