package mediaunlocktest

import "net/http"

func HBOMax(http.Client) Result {
	return Result{Status: StatusNo}
}
