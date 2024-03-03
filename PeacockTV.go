package mediaunlocktest

import (
	"net/http"
	"strings"
)

func PeacockTV(c http.Client) Result {
	resp, err := GET(c, "https://www.peacocktv.com/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	if strings.Contains(resp.Header.Get("location"), "unavailable") {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK}
}
