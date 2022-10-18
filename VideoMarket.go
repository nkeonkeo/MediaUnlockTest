package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func VideoMarket(c http.Client) Result {
	req, err := http.NewRequest("POST", "https://api-p.videomarket.jp/v2/authorize/access_token", strings.NewReader(
		`grant_type=client_credentials&client_id=1eolxdrti3t58m2f2k8yi0kli105743b6f8c8295&client_secret=lco0nndn3l9tcbjdfdwlswmee105743b739cfb5a`,
	))
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("User-Agent", UA_Browser)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	// log.Println(string(b))
	var res resVideoMarketToken
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}

	if res.AccessToken == "" {
		return Result{Success: false}
	}

	req, err = http.NewRequest("POST", "https://api-p.videomarket.jp/v2/api/play/keyissue", strings.NewReader(
		`fullStoryId=118008001&playChromeCastFlag=false&loginFlag=0`,
	))
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("User-Agent", UA_Browser)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Authorization", res.AccessToken)

	resp, err = cdo(c, req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var rpk resVideoMarketPlayKey
	if err := json.Unmarshal(b, &rpk); err != nil {
		return Result{Success: false, Err: err}
	}

	req, err = http.NewRequest("GET", "https://api-p.videomarket.jp/v2/api/play/keyauth?playKey="+rpk.PlayKey+"&deviceType=3&bitRate=0&loginFlag=0&connType=", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("User-Agent", UA_Browser)
	req.Header.Set("X-Authorization", res.AccessToken)

	resp, err = cdo(c, req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	switch resp.StatusCode {
	case 200:
		return Result{Success: true}
	case 408:
		return Result{Success: true}
	case 403:
		return Result{Success: false}
	}
	return Result{Success: false, Info: "failed"}
}

type resVideoMarketToken struct {
	AccessToken string `json:"access_token"`
}

type resVideoMarketPlayKey struct {
	PlayKey string
}
