package main

import (
	"cc-service2/broker"
	"cc-service2/ext"
	"cc-service2/service"
	"cc-service2/storage"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	acs3_accessKey := os.Getenv("CCSERV2_ACS3_ACCESSKEY")
	acs3_secretKey := os.Getenv("CCSERV2_ACS3_SECRETKEY")
	imagestore, err := storage.NewArvanCloudS3("cc-practice-004", "ir-thr-at1",
		"https://s3.ir-thr-at1.arvanstorage.com",
		acs3_accessKey,
		acs3_secretKey)
	if err != nil {
		log.Fatalf("Fatal error at object storage: %v", err)
		os.Exit(1)
	}

	mySQL_username := os.Getenv("CCSERV2_MYSQL_USERNAME")
	mySQL_password := os.Getenv("CCSERV2_MYSQL_PASSWORD")
	mySQL_address := "mysql-container:3306"
	database, err := storage.NewMySQLDB(mySQL_username, mySQL_password, mySQL_address, "ccp1")
	if err != nil {
		log.Fatalf("Fatal error at database: %v", err)
	}

	cloudamq_url := os.Getenv("CCSERV2_AMQP_URL")
	cloudamq := broker.NewCloudAMQ(cloudamq_url, "cc-pr")

	hf := ext.NewHuggingFace(os.Getenv("CCSERV2_HF_APIKEY"))

	srv := service.NewService2(*database, *imagestore, *cloudamq, *hf)
	err = srv.Execute()
	if err != nil {
		log.Fatalf("Could not start service2: %v", err)
	}
}
