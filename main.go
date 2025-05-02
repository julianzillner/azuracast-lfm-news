package main

import (
	"log"
	"os"
	"time"
	_ "time/tzdata"

	"github.com/julianzillner/azuracast-lfm-news/download"
	"github.com/julianzillner/azuracast-lfm-news/upload"
	"github.com/julianzillner/azuracast-lfm-news/utils"
)

func main() {
	location, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		panic(err)
	}

	for {
		now := time.Now().In(location)
		nextRun := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 55, 0, 0, location)

		if now.After(nextRun) {
			nextRun = nextRun.Add(time.Hour)
		}
		duration := nextRun.Sub(now)
		log.Printf("Next Download at: %s (in %v)\n", nextRun.Format("15:04"), duration)
		time.Sleep(duration)

		stationName := os.Getenv("STATION_NAME")
		livePasswort := os.Getenv("LIVE_PASSWORD")
		downloadUrl := "@api.radioadmin.laut.fm/news/1"

		url := "https://" + stationName + ":" + livePasswort + downloadUrl

		err := download.Download(url)
		if err != nil {
			log.Println("Error While downloading:", err)
		} else {
			log.Println("Download Successful!")
			upload.Upload()
			time.Sleep(5 * time.Second)
			utils.DeletFile()
		}
	}
}
