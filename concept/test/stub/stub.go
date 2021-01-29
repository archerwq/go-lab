package stub

import (
	"io/ioutil"
	"time"
)

var configFile = "config.json"

// GetConfig .
func GetConfig() ([]byte, error) {
	return ioutil.ReadFile(configFile)
}

var timeNow = time.Now

// GetDate .
func GetDate() int {
	return timeNow().Day()
}
