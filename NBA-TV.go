package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func NBA_TV(c http.Client) Result {
	resp, err := GET(c, "https://www.nba.com/watch/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "Service is not available in your region") {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK}
}
