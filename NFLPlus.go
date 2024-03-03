package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func NFLPlus(c http.Client) Result {
	url := "https://www.nfl.com/plus/"
	resp, err := GET(c, url)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "nflgamepass") {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK}
}
