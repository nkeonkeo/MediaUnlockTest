package mediaunlocktest

import (
	"net/http"
)

func Popcornflix(c http.Client) Result {
	resp, err := GET(c, "https://popcornflix-prod.cloud.seachange.com/cms/popcornflix/clientconfiguration/versions/2")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		return Result{Status: StatusBanned}
	}
	if resp.StatusCode == 200 {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusFailed}
}
