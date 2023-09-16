package mediaunlocktest

import "net/http"

func ESPNPlus(c http.Client) Result {
	return Result{Status: StatusNo}
}
