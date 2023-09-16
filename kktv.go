package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func KKTV(c http.Client) Result {
	resp, err := GET(c, "https://api.kktv.me/v3/ipcheck")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	// log.Println(string(b))
	var res struct {
		Data struct {
			Country   string
			IsAllowed bool `json:"is_allowed"`
		}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	if res.Data.Country == "TW" && res.Data.IsAllowed {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
