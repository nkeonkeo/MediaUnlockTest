package mediaunlocktest

import "net/http"

func Karaoke(c http.Client) Result {
	resp, err := GET(c, "http://cds1.clubdam.com/vhls-cds1/site/xbox/sample_1.mp4.m3u8")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return Result{Status: StatusOK}
	case 403:
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusUnexpected}
}
