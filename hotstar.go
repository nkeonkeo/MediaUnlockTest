package mediaunlocktest

import (
	"net/http"
	"strings"
)

func Hotstar(c http.Client) Result {
	resp, err := GET(c, "https://api.hotstar.com/o/v1/page/1557?offset=0&size=20&tao=0&tas=20")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 475:
		return Result{Status: StatusNo}
	case 401:
		resp, err := GET(c, "https://www.hotstar.com")
		if err != nil {
			return Result{Status: StatusNetworkErr, Err: err}
		}
		if resp.StatusCode == 301 {
			return Result{Status: StatusNo}
		}
		u := resp.Header.Get("Location")
		if u == "" {
			return Result{Status: StatusNo}
		}
		t := strings.SplitN(u, "/", 4)
		if len(t) < 4 {
			return Result{Status: StatusNo}
		}
		return Result{Status: StatusOK, Region: t[3]}
	}
	return Result{Status: StatusUnexpected}
}
