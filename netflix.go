package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func NetflixRegion(c http.Client) Result {
	// 70143836 绝命毒师
	// 80018499 test
	// 81280792 乐高
	resp, err := GET(c, "https://www.netflix.com/title/81280792")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
	// log.Println(resp.StatusCode)
	if resp.StatusCode == 404 {
		return Result{Success: false, Info: "Originals Only"}
	}
	if resp.StatusCode == 403 {
		return Result{Success: false}
	}
	if resp.StatusCode == 301 || resp.StatusCode == 200 {
		u := resp.Header.Get("location")
		if u == "" {
			return Result{Success: true, Region: "us"}
		}
		// log.Println("nf", u)
		t := strings.SplitN(u, "/", 5)
		if len(t) < 5 {
			return Result{Success: false}
		}
		return Result{Success: true, Region: strings.SplitN(t[3], "-", 2)[0]}
	}
	return Result{Success: false}
}

func NetflixCDN(c http.Client) Result {
	resp, err := GET(c, "https://api.fast.com/netflix/speedtest/v2?https=true&token=YXNkZmFzZGxmbnNkYWZoYXNkZmhrYWxm&urlCount=5")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	if resp.StatusCode == 403 {
		return Result{
			Success: false,
			Info:    "IP Banned By Netflix",
		}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	var res netflixCdnResult
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Success: false, Err: err}
	}
	// u, err := url.Parse(res.Targets[0].Url)
	// if err!=nil{
	// 	return Result{Success: false, Err: err}
	// }
	// ips,err:=net.LookupHost(u.Host)
	// if err!=nil{
	// 	return Result{Success: false, Err: err}
	// }
	return Result{
		Success: true,
		Region:  res.Targets[0].Location.Country,
	}
}

type netflixLocation struct {
	City    string
	Country string
}
type netflixCdnTarget struct {
	Name     string
	Url      string
	Location netflixLocation
}
type netflixCdnResult struct {
	Targets []netflixCdnTarget
}
