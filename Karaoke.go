package mediaunlocktest

import "net/http"

func Karaoke(c http.Client) Result {
	resp, err := GET(c, "http://cds1.clubdam.com/vhls-cds1/site/xbox/sample_1.mp4.m3u8")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return Result{Success: true}
	case 403:
		return Result{Success: false}
	}
	return Result{Success: false, Info: "unknown"}
}
