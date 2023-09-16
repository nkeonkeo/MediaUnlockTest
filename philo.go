package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func Philo(c http.Client) Result {
	resp, err := GET(c, "https://content-us-east-2-fastly-b.www.philo.com/geo")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	if strings.Contains(string(b), "SUCCESS") {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
