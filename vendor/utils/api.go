package utils

import (
	"net/http"
	"encoding/json"
	"errors"
)

// TODO: add inteligently implemented possibility of adjusting the response header
func WriteJSONResponse(w http.ResponseWriter, contents interface{}, status int) error {
		if (http.StatusText(status) == "") { // check if specified error is one of standards-specified ones
			return errors.New("Unknown status code: "+string(status))
		}
		  w.WriteHeader(status)
			  w.Header().Add("Content-Type","application/json")
			  message,err := json.Marshal(contents)
			  if err != nil {
				  http.Error(w, err.Error(), http.StatusInternalServerError)
					  return err
			  }
	  w.Write(message)
  return nil
}
