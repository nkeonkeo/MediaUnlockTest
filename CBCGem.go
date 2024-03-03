package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func CBCGem(c http.Client) Result {
	resp, err := GET(c, "https://www.cbc.ca/g/stats/js/cbc-stats-top.js")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, `country":"CA"`) {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
