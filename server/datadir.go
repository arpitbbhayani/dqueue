package server

import (
	"os"

	"github.com/sirupsen/logrus"
)

func createDataDirectory(path string) {
	err := os.MkdirAll(path, 0664)
	if err != nil {
		logrus.Error(err)
		logrus.Fatalf("unable to create data directory: %s", path)
	}
}
