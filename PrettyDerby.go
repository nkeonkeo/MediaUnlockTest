package mediaunlocktest

import (
	"net/http"
)

func PrettyDerbyJP(c http.Client) Result {
	for i := 0; i < 3; i++ {
		resp, err := GET_Dalvik(c, "https://api-umamusume.cygames.jp/")
		if err != nil {
			return Result{Success: false, Err: err}
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case 404:
			return Result{Success: true}
		}
	}
	return Result{Success: false}
}
