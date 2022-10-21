package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func GYAO(c http.Client) Result {
	resp, err := PostJson(c, "https://gyao.yahoo.co.jp/store/apis/free-playback/graphql?appid=dj00aiZpPWVvSGZ4VmRGbFN0cCZzPWNvbnN1bWVyc2VjcmV0Jng9OTQ-",
		`{"query":" query PlaybackFree($fullStoryId: ID!, $vmDevice: vmDevice!, $isPreview: Boolean) { playback: vmPlaybackFree( fullStoryId: $fullStoryId, isPreview: $isPreview, vmDevice: $vmDevice ) { license { playToken } user { userId } content { encodeVersion maxQualityConsent fullPackId } tracking { streamLog } } } ","variables":{"fullStoryId":"244001026","vmDevice":"PC"}}`,
	)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "foreign") || strings.Contains(s, "DISALLOW_ADDRESS") {
		return Result{Success: false}
	}
	if strings.Contains(s, `"playback":{`) {
		return Result{Success: true}
	}
	return Result{Success: false, Info: "unknown"}
}
