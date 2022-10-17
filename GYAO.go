package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func GYAO(c http.Client) Result {
	req, err := http.NewRequest("POST", "https://gyao.yahoo.co.jp/store/apis/free-playback/graphql?appid=dj00aiZpPWVvSGZ4VmRGbFN0cCZzPWNvbnN1bWVyc2VjcmV0Jng9OTQ-", strings.NewReader(
		`{"query":" query PlaybackFree($fullStoryId: ID!, $vmDevice: vmDevice!, $isPreview: Boolean) { playback: vmPlaybackFree( fullStoryId: $fullStoryId, isPreview: $isPreview, vmDevice: $vmDevice ) { license { playToken } user { userId } content { encodeVersion maxQualityConsent fullPackId } tracking { streamLog } } } ","variables":{"fullStoryId":"244001026","vmDevice":"PC"}}`,
	))
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("user-agent", UA_Browser)
	req.Header.Set("content-type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	s := string(b)
	log.Println(s)
	if strings.Contains(s, "foreign") || strings.Contains(s, "DISALLOW_ADDRESS") {
		return Result{Success: false}
	}
	if strings.Contains(s, `"playback":{`) {
		return Result{Success: true}
	}
	return Result{Success: false, Info: "unknown"}
}
