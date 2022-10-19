package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	m "MediaUnlockTest"

	"github.com/schollz/progressbar/v3"
)

type result struct {
	Name  string
	Value m.Result
}

var R = []*result{}

var bar *progressbar.ProgressBar
var wg = &sync.WaitGroup{}

func excute(Name string, F func(c http.Client) m.Result, C http.Client) {
	r := &result{Name: Name}
	R = append(R, r)
	wg.Add(1)
	go func() {
		r.Value = F(C)
		bar.Describe(Name + " " + ShowResult(r.Value))
		bar.Add(1)
		wg.Done()
	}()
}

var (
	FontRed     = "\033[31m"
	FontGreen   = "\033[32m"
	FontYellow  = "\033[33m"
	FontBlue    = "\033[34m"
	FontPurple  = "\033[35m"
	FontSkyBlue = "\033[36m"
	FontWhite   = "\033[37m"
	FontSuffix  = "\033[0m"
)

func ShowResult(r m.Result) (s string) {
	if r.Success {
		s = FontGreen + "YES"
		if r.Region != "" {
			s += " (Region: " + r.Region + ")"
		}
		s += FontSuffix
	} else {
		if r.Err != nil {
			return FontYellow + "ERR: (" + r.Err.Error() + ")" + FontSuffix
		} else if r.Info != "" {
			return FontRed + " NO" + FontSuffix + FontYellow + " (" + r.Info + ")" + FontSuffix
		} else {
			return FontRed + " NO" + FontSuffix
		}
	}
	return
}

func main() {
	// log.Println(Abema(Ipv4HttpClient))
	// return
	bar = progressbar.Default(48, "testing ...")

	excute("Dazn", m.Dazn, m.Ipv4HttpClient)
	excute("Hotstar", m.Hotstar, m.Ipv4HttpClient)
	excute("Disney+", m.DisneyPlus, m.Ipv4HttpClient)
	excute("Netflix", m.NetflixRegion, m.Ipv4HttpClient)
	excute("Netflix CDN", m.NetflixCDN, m.Ipv4HttpClient)
	excute("Youtube", m.YoutubeRegion, m.Ipv4HttpClient)
	excute("Youtube CDN", m.YoutubeCDN, m.Ipv4HttpClient)
	excute("Prime Video", m.PrimeVideo, m.Ipv4HttpClient)
	excute("TVBAnywhere+", m.TVBAnywhere, m.Ipv4HttpClient)
	excute("iQyi Region", m.IqRegion, m.Ipv4HttpClient)
	excute("Viu.com", m.ViuCom, m.Ipv4HttpClient)
	excute("Spotify", m.Spotify, m.Ipv4HttpClient)
	excute("Steam", m.Steam, m.Ipv4HttpClient)

	excute("Viu.TV", m.ViuTV, m.Ipv4HttpClient)
	excute("Now E", m.NowE, m.Ipv4HttpClient)
	excute("MyTVSuper", m.MyTvSuper, m.Ipv4HttpClient)
	excute("HBO GO Aisa", m.HboGoAisa, m.Ipv4HttpClient)
	excute("BiliBili Hongkong/Macau/Taiwan", m.BilibiliHKMCTW, m.Ipv4HttpClient)

	excute("KKTV", m.KKTV, m.Ipv4HttpClient)
	excute("LiTV", m.LiTV, m.Ipv4HttpClient)
	excute("MyVideo", m.MyVideo, m.Ipv4HttpClient)
	excute("TW4GTV", m.TW4GTV, m.Ipv4HttpClient)
	excute("LineTV", m.LineTV, m.Ipv4HttpClient)
	excute("HamiVideo", m.HamiVideo, m.Ipv4HttpClient)
	excute("Catchplay+", m.Catchplay, m.Ipv4HttpClient)
	excute("Bahamu Anime", m.BahamuAnime, m.Ipv4HttpClient)
	excute("HBO GO Aisa", m.HboGoAisa, m.Ipv4HttpClient)
	excute("Bilibili Taiwan Only", m.BilibiliTW, m.Ipv4HttpClient)

	excute("DMM", m.DMM, m.Ipv4HttpClient)
	excute("Abema", m.Abema, m.Ipv4HttpClient)
	excute("Niconico", m.Niconico, m.Ipv4HttpClient)
	excute("music.jp", m.MusicJP, m.Ipv4HttpClient)
	excute("Telasa", m.Telasa, m.Ipv4HttpClient)
	excute("Paravi", m.Paravi, m.Ipv4HttpClient)
	excute("U-NEXT", m.U_NEXT, m.Ipv4HttpClient)
	excute("Hulu Japan", m.HuluJP, m.Ipv4HttpClient)
	excute("GYAO!", m.GYAO, m.Ipv4HttpClient)
	excute("VideoMarket", m.VideoMarket, m.Ipv4HttpClient)
	excute("FOD(Fuji TV)", m.FOD, m.Ipv4HttpClient)
	excute("Radiko", m.Radiko, m.Ipv4HttpClient)
	excute("Karaoke@DAM", m.Karaoke, m.Ipv4HttpClient)
	excute("J:COM On Demand", m.J_COM_ON_DEMAND, m.Ipv4HttpClient)
	excute("Kancolle", m.Kancolle, m.Ipv4HttpClient)
	excute("Pretty Derby Japan", m.PrettyDerbyJP, m.Ipv4HttpClient)
	excute("Konosuba Fantastic Days", m.KonosubaFD, m.Ipv4HttpClient)
	excute("Princess Connect Re:Dive Japan", m.PCRJP, m.Ipv4HttpClient)
	excute("World Flipper Japan", m.WFJP, m.Ipv4HttpClient)
	excute("Project Sekai: Colorful Stage", m.PJSK, m.Ipv4HttpClient)

	wg.Wait()
	bar.Describe("Finished")
	bar.Clear()
	bar.Finish()

	NameLength := 0
	for _, r := range R {
		if len(r.Name) > NameLength {
			NameLength = len(r.Name)
		}
	}
	for _, r := range R {
		fmt.Printf("%-"+strconv.Itoa(NameLength)+"s %s\n", r.Name, ShowResult(r.Value))
	}
}
