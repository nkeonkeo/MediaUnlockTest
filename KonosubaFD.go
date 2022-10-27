package mediaunlocktest

import "net/http"

func KonosubaFD(c http.Client) Result {
	req, err := http.NewRequest("POST", "https://api.konosubafd.jp/api/masterlist", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("User-Agent", "pj0007/212 CFNetwork/1240.0.4 Darwin/20.6.0")

	resp, err := cdo(c, req)
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
