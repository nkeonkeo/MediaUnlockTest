package mediaunlocktest

import "net/http"

func SlingTV(c http.Client) Result {
	resp, err := GET(c, "https://www.sling.com/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		return Result{Status: StatusNo}
	}
	if resp.StatusCode == 200 {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusFailed}
}
