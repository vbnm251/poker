package handler

import (
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, reason error) {
	log.Println("Error:", reason)
	http.Error(w, reason.Error(), status)
}
