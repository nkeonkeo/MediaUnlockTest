package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func bilibili(c http.Client, u string) Result {
	req, err := http.NewRequest("GET", "https://api.bilibili.com/pgc/player/web/playurl?avid=18281381&cid=29892777&qn=0&type=&otype=json&ep_id=183799&fourk=1&fnver=0&fnval=16&module=bangumi", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Add("user-agent", UA_Browser)

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res bilibiliRes
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.Code == 0 {
		return Result{Success: true}
	}
	if res.Code == -10403 {
		return Result{Success: false}
	}
	return Result{Success: false, Info: "unknown"}
}

func BilibiliHKMCTW(c http.Client) Result {
	return bilibili(c, "https://api.bilibili.com/pgc/player/web/playurl?avid=18281381&cid=29892777&qn=0&type=&otype=json&ep_id=183799&fourk=1&fnver=0&fnval=16&module=bangumi")
}

func BilibiliTW(c http.Client) Result {
	return bilibili(c, "https://api.bilibili.com/pgc/player/web/playurl?avid=50762638&cid=100279344&qn=0&type=&otype=json&ep_id=268176&fourk=1&fnver=0&fnval=16&module=bangumi")
}

type bilibiliRes struct {
	Code int
}
