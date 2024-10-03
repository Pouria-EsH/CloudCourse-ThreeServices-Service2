package service

import (
	"cc-service2/broker"
	"cc-service2/ext"
	"cc-service2/storage"
)

type Service2 struct {
	DataBase  storage.RequestDB
	PicStore  storage.ImageStorage
	MsgBroker broker.CloudAMQ
	ImgDesc   ext.ImgDescriptionService
}

func NewService2(db storage.RequestDB,
	imgstore storage.ImageStorage,
	msgbroker broker.CloudAMQ,
	imgdesc ext.ImgDescriptionService) *Service2 {

	return &Service2{
		DataBase:  db,
		PicStore:  imgstore,
		MsgBroker: msgbroker,
		ImgDesc:   imgdesc,
	}
}

func (s Service2) Execute() error {
	return s.MsgBroker.Listen(s.messageHandler)
}
