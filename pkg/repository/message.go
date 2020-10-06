package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type messageRepository struct {
	urlStr string
}

func NewMessageRepository(url string) Message {
	return &messageRepository{urlStr: url}
}

func (r messageRepository) SendSMS(phoneNumber string, message string) {

	msgData := url.Values{}
	msgData.Set("To", phoneNumber)
	msgData.Set("From", os.Getenv("TWILIO_NUMBER"))
	msgData.Set("Body", message)

	msgReader := strings.NewReader(msgData.Encode())
	client := http.Client{Timeout: 5 * time.Second}

	req, _ := http.NewRequest(http.MethodPost, r.urlStr, msgReader)
	req.SetBasicAuth(os.Getenv("TWILIO_SID"), os.Getenv("TWILIO_AUTH_TOKEN"))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}

}
