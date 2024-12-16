package token

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SalesforceTokenProvider struct {
	ClientID     string
	ClientSecret string
	AuthURL      string
}

func (s *SalesforceTokenProvider) GetAccessToken() (string, error) {
	data := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     s.ClientID,
		"client_secret": s.ClientSecret,
	}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(s.AuthURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get access token: %v", resp.Status)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var response map[string]string
	json.Unmarshal(body, &response)

	return response["access_token"], nil
}
