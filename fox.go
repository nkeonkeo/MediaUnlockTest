package mediaunlocktest

import "net/http"

func Fox(c http.Client) Result {
	url := "https://x-live-fox-stgec.uplynk.com/ausw/slices/8d1/d8e6eec26bf544f084bad49a7fa2eac5/8d1de292bcc943a6b886d029e6c0dc87/G00000000.ts?pbs=c61e60ee63ce43359679fb9f65d21564&cloud=aws&si=0"
	resp, err := GET(c, url)
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
