package main

import "net/http"

func J_COM_ON_DEMAND(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://id.zaq.ne.jp", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("User-Agent", UA_Browser)

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	switch resp.StatusCode {
	case 404:
		return Result{Success: true}
	case 403:
		return Result{Success: false}
	}
	return Result{Success: false, Info: "unknown"}
}
