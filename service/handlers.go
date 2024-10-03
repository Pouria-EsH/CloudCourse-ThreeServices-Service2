package service

import "fmt"

func (s Service2) messageHandler(msg string) error {
	fmt.Printf("Received a new request: %s\n", msg)
	reqId := msg

	fmt.Println("downloading...")
	imgFile, err := s.PicStore.Download(reqId)
	if err != nil || imgFile == nil {
		return fmt.Errorf("error at image download: %w", err)
	}

	caption, err := s.ImgDesc.GetDiscription(imgFile)
	if err != nil {
		return fmt.Errorf("error at image description service: %w", err)
	}

	err = s.DataBase.SetImageCaption(reqId, caption)
	if err != nil {
		return fmt.Errorf("error updating image caption: %w", err)
	}

	err = s.DataBase.SetStatus(reqId, "ready")
	if err != nil {
		return fmt.Errorf("error updating request status: %w", err)
	}
	return nil
}
