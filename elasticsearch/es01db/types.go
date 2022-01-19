package es01db

import (
	"encoding/json"
	"log"
	"os"
)

type contextKey struct {
	Key int
}

// MoviesKey ...
var MoviesKey contextKey = contextKey{Key: 1}

// ClientKey ...
var ClientKey contextKey = contextKey{Key: 2}

// ReadJSONFile reads filename fName and unmarshall json to data
func ReadJSONFile(fName string, data interface{}) {
	dataByte, err := os.ReadFile(fName)
	if err != nil {

		return
	}
	err = json.Unmarshal(dataByte, data)
	if err != nil {
		log.Fatalf("[ oi ] couldnt unmarshal data to json. error [%s]", err)
	}
}
