package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func TW4GTV(c http.Client) Result {
	// local tmpresult=$(curl $useNIC $xForward --user-agent "${UA_Browser}" -${1} ${ssll} -sS --max-time 10 -X POST -d  " 2>&1)
	req, err := http.NewRequest("POST", "https://api2.4gtv.tv//Vod/GetVodUrl3", strings.NewReader(
		`value=D33jXJ0JVFkBqV%2BZSi1mhPltbejAbPYbDnyI9hmfqjKaQwRQdj7ZKZRAdb16%2FRUrE8vGXLFfNKBLKJv%2BfDSiD%2BZJlUa5Msps2P4IWuTrUP1%2BCnS255YfRadf%2BKLUhIPj`,
	))
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res struct {
		Success bool
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	if res.Success {
		return Result{Success: true}
	}
	return Result{Success: false}
}
