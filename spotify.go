package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func Spotify(c http.Client) Result {
	req, err := http.NewRequest("POST", "https://spclient.wg.spotify.com/signup/public/v1/account", strings.NewReader(
		`birth_day=11&birth_month=11&birth_year=2000&collect_personal_info=undefined&creation_flow=&creation_point=https%3A%2F%2Fwww.spotify.com%2Fhk-en%2F&displayname=Gay%20Lord&gender=male&iagree=1&key=a1e486e2729f46d6bb368d6b2bcda326&platform=www&referrer=&send-email=0&thirdpartyemail=0&identifier_token=AgE6YTvEzkReHNfJpO114514`,
	))
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("User-Agent", UA_Browser)
	resp, err := cdo(c, req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res spotifyRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.Status == 320 {
		return Result{Success: false}
	}
	if res.Status == 311 && res.IsCountryLaunched {
		return Result{Success: true, Region: strings.ToLower(res.Country)}
	}
	return Result{Success: false}
}

type spotifyRes struct {
	Status            int
	Country           string
	IsCountryLaunched bool `json:"is_country_launched"`
}
