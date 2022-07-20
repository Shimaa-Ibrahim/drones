package apis

import (
	"fmt"
	"log"
	"net/http"
)

func sendError(w http.ResponseWriter, statusCode int, err error) {
	log.Printf("[Error]: %v", err.Error())
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"error":%q}`, err.Error())))
}
