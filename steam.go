package mediaunlocktest

import (
	"net/http"
	"strings"
)

func Steam(c http.Client) Result {
	resp, err := GET(c, "https://store.steampowered.com")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	for _, c := range resp.Cookies() {
		if c.Name == "steamCountry" {
			i := strings.Index(c.Value, "%")
			if i == -1 {
				return Result{Status: StatusNo}
			}
			return Result{Status: StatusOK, Region: strings.ToLower(c.Value[:i])}
		}
	}
	return Result{Status: StatusNo}
}
