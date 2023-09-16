package mediaunlocktest

import "net/http"

func EncoreTVB(c http.Client) Result {
	return Result{Status: StatusNo}
}
