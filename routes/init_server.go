package routes

import (
	"encoding/json"
	"net/http"

	"github.com/emanueldias01/TranscribeAudiosAPI/model"
	"github.com/emanueldias01/TranscribeAudiosAPI/service"
)


func InitServer(){
	http.HandleFunc("/transcription", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        text, err := service.TranscribeAudio(r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

		resp := model.Transcription{Transcription: text}

        json.NewEncoder(w).Encode(resp)

	})
	http.ListenAndServe(":8080", nil)
}
