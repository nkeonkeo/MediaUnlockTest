package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func MyTvSuper(c http.Client) Result {
	r, err := http.NewRequest("GET", "https://www.mytvsuper.com/api/auth/getSession/self/", nil)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	r.Header.Set("User-Agent", UA_Browser)
	r.Header.Set("Content-Type", "application/json")

	resp, err := cdo(c, r)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	var res mytvsuperRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	if res.Region == 1 {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}

type mytvsuperRes struct {
	Region int
}
