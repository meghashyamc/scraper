package scraperapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func writeResponse(w *http.ResponseWriter, object interface{}, status string, statusCode int) {
	resp := apiResponse{Status: status, Data: object}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Println("error in writing response:", err)
		return
	}
	fmt.Println("respBytes", string(respBytes))
	(*w).WriteHeader(statusCode)
	(*w).Write(respBytes)
}
