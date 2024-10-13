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
	if err := godotenv.Load(".env"); err != nil {
		log.Println("no .env file found")
	}

	acs3_bucket := os.Getenv("CCSERV1_ACS3_BUCKET")
	acs3_region := os.Getenv("CCSERV1_ACS3_REGION")
	acs3_endpoint := os.Getenv("CCSERV1_ACS3_ENDPOINT")
	acs3_accessKey := os.Getenv("CCSERV2_ACS3_ACCESSKEY")
	acs3_secretKey := os.Getenv("CCSERV2_ACS3_SECRETKEY")
	imagestore, err := storage.NewArvanCloudS3(
		acs3_bucket,
		acs3_region,
		acs3_endpoint,
		acs3_accessKey,
		acs3_secretKey)
	if err != nil {
		log.Fatalf("Fatal error at object storage: %v\n", err)
	}

	mySQL_username := os.Getenv("CCSERV2_MYSQL_USERNAME")
	mySQL_password := os.Getenv("CCSERV2_MYSQL_PASSWORD")
	mySQL_address := "mysql-container:3306"
	database, err := storage.NewMySQLDB(mySQL_username, mySQL_password, mySQL_address, "ccp1")
	if err != nil {
		log.Fatalf("Fatal error at database: %v\n", err)
	}

	cloudamq_url := os.Getenv("CCSERV2_AMQP_URL")
	cloudamq := broker.NewCloudAMQ(cloudamq_url, "cc-pr")

	hf := ext.NewHuggingFace(os.Getenv("CCSERV2_HF_APIKEY"))

	srv := service.NewService2(*database, *imagestore, *cloudamq, *hf)
	err = srv.Execute()
	if err != nil {
		log.Fatalf("Could not start service2: %v\n", err)
	}
}
