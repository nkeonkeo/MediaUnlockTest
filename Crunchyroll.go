package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func Crunchyroll(c http.Client) Result {
	resp, err := GET(c, "https://c.evidon.com/geo/country.js")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "'code':'us'") {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
