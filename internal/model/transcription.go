package model

import "fmt"

type Transcription struct {
	IdFile   string `json:"id_file"`
	Output   string `json:"output"`
	Encode   bool   `json:"encode"`
	Language string `json:"language"`
}

func DefaultTranscription(id string) *Transcription {
	return &Transcription{
		IdFile:   id,
		Output:   "json",
		Encode:   true,
		Language: "ru",
	}
}

func (t *Transcription) String() string {
	return fmt.Sprintf("http://localhost:9000/asr?encode=%t"+
		"&task=transcribe&"+
		"language=%s&word_timestamps=false&"+
		"output=%s", t.Encode, t.Language, t.Output)
}
