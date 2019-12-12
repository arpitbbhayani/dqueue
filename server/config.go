package server

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const defaultConfig = `
# port defines the port on which the Dqueue server should work on.
port: 4096

data_dir: /etc/dqueue/data
`

func createDefaultConfigFile(path string) {
	parentDirectory := filepath.Dir(path)

	err := os.MkdirAll(parentDirectory, 0664)
	if err != nil {
		logrus.Error(err)
		logrus.Fatalf("unable to create default configuration file: %s", path)
	}

	err = ioutil.WriteFile(path, []byte(defaultConfig), 0644)
	if err != nil {
		logrus.Error("1", err)
		logrus.Fatalf("unable to create default configuration file: %s", path)
	}
}

func initializeConfig(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logrus.Warnf("configuration file does not exists at %s hence creating default config", path)
		createDefaultConfigFile(path)
	}

	logrus.Infof("reading config file %s", path)
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Fatalf("unable to read configuration file: %s", path)
	}

	viper.SetConfigType("yml")

	// Default values to configurations
	viper.SetDefault("port", 4096)
	viper.SetDefault("data_dir", "/etc/dqueue/data")

	err = viper.ReadConfig(bytes.NewReader(configBytes))
	if err != nil {
		logrus.Fatalf("unable to read configuration file: %s", path)
	}
}
