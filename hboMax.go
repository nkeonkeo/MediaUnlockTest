package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func HBOMax(c http.Client) Result {
	resp, err := GET(c, "https://www.hbomax.com/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "geo-availability") {
		return Result{Status: StatusNo}
	}
	t := strings.Split(resp.Header.Get("location"), "/")
	region := ""
	if len(t) >= 4 {
		region = strings.Split(resp.Header.Get("location"), "/")[3]
	}
	return Result{Status: StatusOK, Region: strings.ToUpper(region)}
}
