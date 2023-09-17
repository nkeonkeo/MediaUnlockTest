package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func VideoMarket(c http.Client) Result {
	resp, err := PostForm(c, "https://api-p.videomarket.jp/v2/authorize/access_token",
		`grant_type=client_credentials&client_id=1eolxdrti3t58m2f2k8yi0kli105743b6f8c8295&client_secret=lco0nndn3l9tcbjdfdwlswmee105743b739cfb5a`,
	)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	// log.Println(string(b))
	var res resVideoMarketToken
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}

	if res.AccessToken == "" {
		return Result{Status: StatusNo}
	}

	req, err := http.NewRequest("POST", "https://api-p.videomarket.jp/v2/api/play/keyissue", strings.NewReader(
		`fullStoryId=118008001&playChromeCastFlag=false&loginFlag=0`,
	))
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	req.Header.Set("User-Agent", UA_Browser)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Authorization", res.AccessToken)

	resp, err = cdo(c, req)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	var rpk resVideoMarketPlayKey
	if err := json.Unmarshal(b, &rpk); err != nil {
		return Result{Status: StatusErr, Err: err}
	}

	req, err = http.NewRequest("GET", "https://api-p.videomarket.jp/v2/api/play/keyauth?playKey="+rpk.PlayKey+"&deviceType=3&bitRate=0&loginFlag=0&connType=", nil)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	req.Header.Set("User-Agent", UA_Browser)
	req.Header.Set("X-Authorization", res.AccessToken)

	resp, err = cdo(c, req)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return Result{Status: StatusOK}
	case 408:
		return Result{Status: StatusOK}
	case 403:
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusUnexpected}
}

type resVideoMarketToken struct {
	AccessToken string `json:"access_token"`
}

type resVideoMarketPlayKey struct {
	PlayKey string
}
