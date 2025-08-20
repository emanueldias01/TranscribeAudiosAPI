package routes

import (
	"encoding/json"
	"net/http"
)


func InitServer(){
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		

		respPong := struct{
			Message string `json:"message"`
		}{
			Message: "pong",
		}
		json.NewEncoder(w).Encode(respPong)
	})
	http.ListenAndServe(":8080", nil)
}