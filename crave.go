package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func Crave(c http.Client) Result {
	resp, err := GET(c, "https://capi.9c9media.com/destinations/se_atexace/platforms/desktop/bond/contents/2205173/contentpackages/4279732/manifest.mpd")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, `Geo Constraint Restrictions`) {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusOK}
}
