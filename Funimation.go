package mediaunlocktest

import (
	"net/http"
)

func Funimation(c http.Client) Result {
	resp, err := GET(c, "https://www.funimation.com")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		return Result{Status: StatusNo}
	}
	for _, c := range resp.Cookies() {
		if c.Name == "region" {
			return Result{Status: StatusOK, Region: c.Value}
		}
	}
	return Result{Status: StatusFailed}
}
