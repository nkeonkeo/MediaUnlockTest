package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func BahamutAnime(c http.Client) Result {
	resp, err := GET(c, "https://ani.gamer.com.tw/ajax/token.php?adID=89422&sn=14667")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res struct {
		AnimeSn int
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false}
	}
	if res.AnimeSn != 0 {
		return Result{Success: true}
	}
	return Result{Success: false}
}
