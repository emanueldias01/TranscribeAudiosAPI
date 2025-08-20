package routes

import (
	"encoding/json"
	"net/http"

	"github.com/emanueldias01/TranscribeAudiosAPI/service"
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


	http.HandleFunc("/transcrible", func(w http.ResponseWriter, r *http.Request) {
		response, err := service.TranscribeAudio(r)

		if err != nil{
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body := struct{
			Response string `json:"response"`
		}{
			Response: response,
		}

		json.NewEncoder(w).Encode(body)
	})
	http.ListenAndServe(":8080", nil)
}
