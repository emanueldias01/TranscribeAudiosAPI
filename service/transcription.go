package service

import (
	"net/http"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func TranscribeAudio(r *http.Request) (string, error) {

	apiKey := os.Getenv("OPENAI_API_KEY")
    client := openai.NewClient(option.WithAPIKey(apiKey))

    resp, err := client.Audio.Transcriptions.New(r.Context(), openai.AudioTranscriptionNewParams{
        File:  r.Body,
        Model: openai.AudioModelWhisper1,
        ResponseFormat: openai.AudioResponseFormatJSON,
        Language: openai.String("pt"),
    })
    if err != nil {
        return "", err
    }

    return resp.Text, nil
}