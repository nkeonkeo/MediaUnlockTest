package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func DMM(c http.Client) Result {
	resp, err := GET(c, "https://bitcoin.dmm.com")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "This page is not available in your area") {
		return Result{Status: StatusNo}
	}
	if strings.Contains(s, "暗号資産") {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo, Info: "Unsupported"}
}
func DMMTV(c http.Client) Result {
	resp, err := PostJson(c, "https://api.beacon.dmm.com/v1/streaming/start", `{"player_name":"dmmtv_browser","player_version":"0.0.0","content_type_detail":"VOD_SVOD","content_id":"11uvjcm4fw2wdu7drtd1epnvz","purchase_product_id":null}`)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "FOREIGN") {
		return Result{Status: StatusNo}
	}
	if strings.Contains(s, "UNAUTHORIZED") {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo, Info: "Unsupported"}
}
