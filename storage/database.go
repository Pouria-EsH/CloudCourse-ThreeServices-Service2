package storage

import "fmt"

type RequestDB interface {
	SetStatus(requestId string, status string) error
	SetImageCaption(requestId string, imgcap string) error
}

type RequestNotFoundError struct {
	ReqId string
}

func (e RequestNotFoundError) Error() string {
	return fmt.Sprintf("request with id %s not found", e.ReqId)
}

type PicRequestEntry struct {
	ReqId        string
	Email        string
	ReqStatus    string
	ImageCaption string
	NewImageURL  string
}
