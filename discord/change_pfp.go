package discord

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

type Avatar struct {
	Avatar string `json:"avatar"`
}

// The ChangePFP function changes the profile picture of the user. The image must be a base64 encoded string.
func ChangePFP(token, image string) error {

	var avatar Avatar
	avatar.Avatar = "data:image/png;base64," + image
	jsonBytes, err := json.Marshal(avatar)
	if err != nil {
		return errors.New("error marshalling json: " + err.Error())
	}
	jsonStr := string(jsonBytes)

	//log.Println(jsonStr)

	url := "https://discord.com/api/v9/users/@me"
	req, _ := http.NewRequest("PATCH", url, strings.NewReader(jsonStr))

	for k, v := range discordCommonHeaders(token) {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("error sending request: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error reading response: " + err.Error())
	}
	log.Println(string(body))

	return nil
}

// The discordCommonHeaders function returns a map of common headers used in Discord requests.
func discordCommonHeaders(authorization string) map[string]string {
	return map[string]string{
		"accept":             "*/*",
		"accept-language":    "en-US",
		"authorization":      authorization,
		"content-type":       "application/json",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDEwIiwib3NfdmVyc2lvbiI6IjEwLjAuMjI2MjEiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTczNTAxLCJuYXRpdmVfYnVpbGRfbnVtYmVyIjoyOTEyOCwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=",
	}
}

type RatelimitResponse struct {
	Global     bool    `json:"global"`
	Message    string  `json:"message"`
	RetryAfter float64 `json:"retry_after"`
}
