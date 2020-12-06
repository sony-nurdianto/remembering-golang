package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sony-nurdianto/remembering-golang/models"
)

//SendError status message
func SendError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

//SendSucces status message
func SendSucces(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

//LogFatal simple handle error
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
