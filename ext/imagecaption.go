package ext

import "bytes"

type ImgDescriptionService interface {
	GetDiscription(fileBytes *bytes.Buffer) (string, error)
}
