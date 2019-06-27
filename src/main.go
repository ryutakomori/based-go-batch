package main

import (
	"fmt"
	"os"
	"time"
	log "util/log"
)

const (
	ARG_FUNCTION = 1
)

func main() {
	log.Intialize()

	startTime := time.Now()
	fmt.Println(startTime)

	fn := os.Args[ARG_FUNCTION]

	//slack 開始時間通知
	text := fn + " start at " + startTime.Format("2006-01-02 15:04:05")
	log.Println(text)

	//slack.Slack(text)

	log.Println("batch " + fn + " start")

	switch fn {
	// case "HoujinCSV":
	// 	bt.HoujinCSV()

	// case "HoujinAPI":
	// 	bt.HoujinAPI()

	case "test":
		log.Println("テスト")

	default:

	}
	//slack 終了時間通知
	elapsedTime := time.Now().Sub(startTime)
	text = fmt.Sprintf("%s elapsed time %f [sec]", fn, elapsedTime.Minutes())
	//slack.Slack(text)
	log.Println(text)
	log.Println("batch " + fn + " end")

}
