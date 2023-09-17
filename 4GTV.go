package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func TW4GTV(c http.Client) Result {
	resp, err := PostForm(c, "https://api2.4gtv.tv//Vod/GetVodUrl3",
		`value=D33jXJ0JVFkBqV%2BZSi1mhPltbejAbPYbDnyI9hmfqjKaQwRQdj7ZKZRAdb16%2FRUrE8vGXLFfNKBLKJv%2BfDSiD%2BZJlUa5Msps2P4IWuTrUP1%2BCnS255YfRadf%2BKLUhIPj`,
	)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	var res struct {
		Success bool
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	if res.Success {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}
