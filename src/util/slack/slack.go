package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var (
	// slackから取得したURL
	IncomingUrl string = ""
)

type slack struct {
	Text       string `json:"text"`
	Username   string `json:"username"`
	Icon_emoji string `json:"icon_emoji"`
	Icon_url   string `json:"icon_url"`
	Channel    string `json:"channel"`
}

func Slack(text string) {
	mode := os.Getenv("MODE")

	params, _ := json.Marshal(slack{
		fmt.Sprintf("%s %s", mode, text),
		"batch-bot",
		"",
		"",
		"#batch"})

	resp, _ := http.PostForm(
		IncomingUrl,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
}
