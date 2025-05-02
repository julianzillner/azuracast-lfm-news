package utils

import (
	"log"
	"os"
)

func DeletFile() {
	filePath := "news.mp3"
	errs := os.Remove(filePath)
	if errs != nil {
		log.Fatalf("Error removing file: %v", errs)
	}
	log.Printf("File %s removed successfully", filePath)
}
