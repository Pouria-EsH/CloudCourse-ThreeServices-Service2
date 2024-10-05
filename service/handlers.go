package service

import (
	"cc-service2/storage"
	"errors"
	"fmt"
)

func (s Service2) messageHandler(requestID string) error {
	fmt.Printf("Received a new request: %s\n", requestID)

	imgFile, err := s.PicStore.Download(requestID)
	if err != nil || imgFile == nil {
		s.failureHandler(requestID)
		return fmt.Errorf("error at image download: %w", err)
	}

	caption, err := s.ImgDesc.GetDiscription(imgFile)
	if err != nil {
		s.failureHandler(requestID)
		return fmt.Errorf("error at image description service: %w", err)
	}

	err = s.DataBase.SetImageCaption(requestID, caption)
	if err != nil {
		return fmt.Errorf("error updating image caption: %w", err)
	}

	err = s.DataBase.SetStatus(requestID, "ready")
	if err != nil {
		return fmt.Errorf("error updating request status: %w", err)
	}
	fmt.Printf("Request %s done successfuly\n", requestID)
	return nil
}

func (s Service2) failureHandler(requstId string) {
	err := s.DataBase.SetStatus(requstId, "failure")
	if err != nil {
		var notfound *storage.RequestNotFoundError
		if !errors.As(err, &notfound) {
			fmt.Printf("couldn't update request %s status to \"failed\": %v\n", requstId, err)
		}
		return
	}
	fmt.Printf("request %s status is set to 'failure'", requstId)
}
