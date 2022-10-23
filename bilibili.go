package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func bilibili(c http.Client, url string) Result {
	resp, err := GET(c, url)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res struct {
		Code int
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.Code == -10403 {
		return Result{Success: false}
	}
	if res.Code == 0 {
		return Result{Success: true}
	}
	return Result{Success: false, Info: "unknown"}
}

func BilibiliHKMC(c http.Client) Result {
	return bilibili(c, "https://api.bilibili.com/pgc/player/web/playurl?avid=473502608&cid=845838026&qn=0&type=&otype=json&ep_id=678506&fourk=1&fnver=0&fnval=16&module=bangumi")
}

func BilibiliTW(c http.Client) Result {
	return bilibili(c, "https://api.bilibili.com/pgc/player/web/playurl?avid=50762638&cid=100279344&qn=0&type=&otype=json&ep_id=268176&fourk=1&fnver=0&fnval=16&module=bangumi")
}
