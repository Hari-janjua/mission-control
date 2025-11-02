package http

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type AuthClient struct {
	BaseURL    string
	SoldierID  string
	Token      string
	ExpiryTime time.Time
}

func NewAuthClient(baseURL, soldierID string) *AuthClient {
	return &AuthClient{BaseURL: baseURL, SoldierID: soldierID}
}

func (a *AuthClient) GetToken() (string, error) {
	if time.Now().Before(a.ExpiryTime.Add(-5*time.Second)) && a.Token != "" {
		return a.Token, nil
	}

	body, _ := json.Marshal(map[string]string{"soldier_id": a.SoldierID})
	resp, err := http.Post(a.BaseURL+"/auth/token", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res struct {
		Token string `json:"token"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&res)

	a.Token = res.Token
	a.ExpiryTime = time.Now().Add(30 * time.Second)
	log.Printf("[Auth] Refreshed token for %s", a.SoldierID)
	return a.Token, nil
}
