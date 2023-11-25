package service

import (
	"bytes"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
	"time"
	"transcription/internal/model"
)

type transcriptionService struct {
	AS string
}

func (t *transcriptionService) Default(id string) ([]byte, error) {
	res, err := http.Get(t.AS + id)
	if err != nil {
		return nil, errors.Wrap(err, "Trans/Default: fail GET request to AS")
	}

	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	fw, err := writer.CreateFormFile("audio_file", id+".mp3")
	if err != nil {
		return nil, errors.Wrap(err, "Trans/Default: fail create formfile")
	}
	_, err = io.Copy(fw, res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Trans/Default: fail copy")
	}
	writer.Close()

	defauluri := model.DefaultTranscription(id).String()
	client := &http.Client{Timeout: 1 * time.Minute}
	req, err := http.NewRequest("POST", defauluri, form)
	if err != nil {
		return nil, errors.Wrap(err, "Trans/Default: fail copy")
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Trans/Default: fail request")
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Trans/Default: fail to read response body")
	}
	return bodyText, nil
}

func (t *transcriptionService) Custom(trans *model.Transcription) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func newTranscriptionService(AS string) *transcriptionService {
	return &transcriptionService{AS: AS}
}
