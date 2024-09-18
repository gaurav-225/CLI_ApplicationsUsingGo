package models

import (
	"bytes"
	"encoding/json"
)

type Audio struct {
	Id string `json:"Id"`
	Path string `json:"Path"`
	Metadata Metadata `json:"Metadata"`
	Status string `json:"Status"`
	Error []string `json:"Error"`
}


func (a *Audio)	JSON() (string, error){
	// Mashal serializes the data into a JSON byte slice
	jsonData, err := json.Marshal(a)
	if err != nil {
		return "", err
	}


	// bytes.BUffer is used to build and store byte slices
	// we are using it to format the JSON data
	var prettyJSON bytes.Buffer

	if err := json.Indent(&prettyJSON, jsonData, "", "\t"); err != nil {
		return "", err
	}

	// prettyJSON.String() converts the bytes.Buffer to a string
	return prettyJSON.String(), nil
}