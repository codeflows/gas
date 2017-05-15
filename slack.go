package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// SendAdToSlack formats and sends and ad to Slack
func SendAdToSlack(ad Ad) (err error) {
	text := fmt.Sprintf("%s: <%s|%s> %s", ad.category, ad.url, ad.title, ad.price)

	payload := map[string]string{"text": text}
	jsonPayload, jsonError := json.Marshal(payload)
	if jsonError != nil {
		return jsonError
	}

	resp, err := http.PostForm(
		os.Getenv("SLACK_WEBHOOK_URL"),
		url.Values{"payload": {string(jsonPayload)}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Slack returned non-200 status code: %d", resp.StatusCode)
	}
	return nil
}
