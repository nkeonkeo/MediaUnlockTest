package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func U_NEXT(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://video-api.unext.jp/api/1/player?entity%5B%5D=playlist_url&episode_code=ED00148814&title_code=SID0028118&keyonly_flg=0&play_mode=caption&bitrate_low=1500", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("User-Agent", UA_Browser)

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	// log.Println(string(b))
	var res struct {
		Data struct {
			EntitiesData struct {
				PlaylistUrl struct {
					ResultStatus int `json:"result_status"`
				} `json:"playlist_url"`
			} `json:"entities_data"`
		}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	switch res.Data.EntitiesData.PlaylistUrl.ResultStatus {
	case 475:
		return Result{Success: true}
	case 200:
		return Result{Success: true}
	case 467:
		return Result{Success: false}
	}
	return Result{Success: false, Info: "unknown"}
}
