package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"transcription/internal/service"
)

type transcriptionController struct {
	T service.Transcription
}

func newTranscriptionController(t service.Transcription) *transcriptionController {
	return &transcriptionController{T: t}
}

func NewTranscriptionRoutes(r *mux.Router, t service.Transcription) {
	ctrl := newTranscriptionController(t)

	//r.HandleFunc("/transcription/{id:[0-9]+}", ctrl).Methods(http.MethodPost)
	r.HandleFunc("/transcription/{id:[0-9]+}", ctrl.defaultTranscription).Methods(http.MethodGet)
}

func (t *transcriptionController) defaultTranscription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bytes, err := t.T.Default(vars["id"])
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
	}
}
