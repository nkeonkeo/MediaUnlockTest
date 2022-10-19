package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func NetflixRegion(c http.Client) Result {
	r, err := http.NewRequest("GET", "https://www.netflix.com/title/80018499", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	r.Header.Set("User-Agent", UA_Browser)

	resp, err := c.Do(r)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return Result{Success: false, Info: "Originals Only"}
	}
	if resp.StatusCode == 403 {
		return Result{Success: false}
	}
	if resp.StatusCode == 200 {
		return Result{Success: true, Region: "us"}
	}
	if resp.StatusCode == 301 {
		u := resp.Header.Get("location")
		if u == "" {
			return Result{Success: false}
		}
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
