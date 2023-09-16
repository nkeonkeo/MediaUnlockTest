package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
)

func BahamutAnime(c http.Client) Result {
	c.Jar, _ = cookiejar.New(nil)
	resp, err := GET(c, "https://ani.gamer.com.tw/ajax/getdeviceid.php")
	if err != nil {
		return Result{Status: StatusNetworkErr}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	var res struct {
		AnimeSn  int
		Deviceid string
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr}
	}
	resp, err = GET(c, "https://ani.gamer.com.tw/ajax/token.php?adID=89422&sn=14667&device="+res.Deviceid)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr}
	}
	if res.AnimeSn != 0 {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
