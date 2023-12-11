package mediaunlocktest

import "net/http"

func ESPNPlus(c http.Client) Result {
	req, err := http.NewRequest("POST", "https://espn.api.edge.bamgrid.com/token", nil)
	if err != nil {
		return Result{}
	}
	req.Header.Set("authorization", "Bearer ZXNwbiZicm93c2VyJjEuMC4w.ptUt7QxsteaRruuPmGZFaJByOoqKvDP2a5YkInHrc7c")

	// resp, err := cdo(c, req)
	// if err != nil {
	// 	return Result{Status: StatusNetworkErr, Err: err}
	// }

	return Result{Status: StatusNo}
}
