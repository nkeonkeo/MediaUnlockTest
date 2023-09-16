package mediaunlocktest

import (
	"net/http"
	"strings"
)

func ViuCom(c http.Client) Result {
	resp, err := GET(c, "https://www.viu.com")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if location := resp.Header.Get("location"); location != "" {
		region := strings.Split(location, "/")[4]
		if region == "no-service" {
			return Result{Status: StatusNo}
		}
		return Result{Status: StatusOK, Region: region}
	}
	return Result{Status: StatusNo}
}
