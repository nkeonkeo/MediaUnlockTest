package mediaunlocktest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func Epix(c http.Client) Result {
	url := "https://api.epix.com/v2/sessions"
	resp, err := PostJson(c, url, `{"device":{"guid":"e2add88e-2d92-4392-9724-326c2336013b","format":"console","os":"web","app_version":"1.0.2","model":"browser","manufacturer":"google"},"apikey":"f07debfcdf0f442bab197b517a5126ec","oauth":{"token":null}}`)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "error code") {
		return Result{Status: StatusNo}
	}
	log.Print(s)
	var res struct {
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusNo}
}
