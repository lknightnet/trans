package service

import "transcription/internal/model"

type Transcription interface {
	Default(id string) ([]byte, error)
	Custom(trans *model.Transcription) ([]byte, error)
}

type Services struct {
	Transcription
}

type Dependencies struct {
	AS string
}

func NewServices(deps *Dependencies) *Services {
	return &Services{Transcription: newTranscriptionService(deps.AS)}
}
