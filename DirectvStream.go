package mediaunlocktest

import (
	"net/http"
)

func DirectvStream(c http.Client) Result {
	resp, err := GET(c, "https://www.atttvnow.com/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		return Result{Status: StatusNo}
	}
	// if resp.StatusCode == 200 || resp.StatusCode == 301 {
	// 	return Result{Status: StatusOK}
	// }
	return Result{Status: StatusOK}
}
