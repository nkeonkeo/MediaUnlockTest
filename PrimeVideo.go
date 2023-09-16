package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func PrimeVideo(c http.Client) Result {
	c.CheckRedirect = nil
	resp, err := GET(c, "https://www.primevideo.com")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if i := strings.Index(s, `"currentTerritory":`); i != -1 {
		return Result{
			Status: StatusOK,
			Region: strings.ToLower(s[i+20 : i+22]),
		}
	}
	return Result{Status: StatusNo}
}
