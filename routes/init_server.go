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
		if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        text, err := service.TranscribeAudio(r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write([]byte(text))

	})
	http.ListenAndServe(":8080", nil)
}
