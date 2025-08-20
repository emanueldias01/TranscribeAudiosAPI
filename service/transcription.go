package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func TranscribeAudio(r *http.Request) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(option.WithAPIKey(apiKey))

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return "", err
	}

	fileHeader, ok := r.MultipartForm.File["file"]
	if !ok || len(fileHeader) == 0 {
		return "", fmt.Errorf("nenhum arquivo enviado com a chave 'file'")
	}

	fh := fileHeader[0]

	src, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	tempDir := os.TempDir()
	tempFilePath := filepath.Join(tempDir, fh.Filename)

	dst, err := os.Create(tempFilePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	audioFile, err := os.Open(tempFilePath)
	if err != nil {
		return "", err
	}
	defer audioFile.Close()

	resp, err := client.Audio.Transcriptions.New(r.Context(), openai.AudioTranscriptionNewParams{
		File:           audioFile,
		Model:          openai.AudioModelWhisper1,
		ResponseFormat: openai.AudioResponseFormatJSON,
		Language:       openai.String("pt"),
	})
	if err != nil {
		return "", err
	}

	os.Remove(tempFilePath)

	return resp.Text, nil
}
