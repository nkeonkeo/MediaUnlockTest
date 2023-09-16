package main

import (
	mt "MediaUnlockTest"
	"net/http"
	"sync"
)

var R map[string]mt.Result

var tot int64
var wg *sync.WaitGroup

func excute(Name string, F func(client http.Client) mt.Result, client http.Client) {
	if _, has := R[Name]; !has {
		R[Name] = mt.Result{Info: "Waiting"}
	}
	wg.Add(1)
	go func() {
		r := F(client)
		R[Name] = r
		wg.Done()
	}()
}

func Multination(c http.Client) {
	// R = append(R, &result{Name: "Multination", Divider: true})
	excute("Dazn", mt.Dazn, c)
	excute("Hotstar", mt.Hotstar, c)
	excute("Disney+", mt.DisneyPlus, c)
	excute("Netflix", mt.NetflixRegion, c)
	excute("Netflix CDN", mt.NetflixCDN, c)
	excute("Youtube", mt.YoutubeRegion, c)
	excute("Youtube CDN", mt.YoutubeCDN, c)
	excute("Amazon Prime Video", mt.PrimeVideo, c)
	excute("TVBAnywhere+", mt.TVBAnywhere, c)
	excute("iQyi", mt.IqRegion, c)
	excute("Viu.com", mt.ViuCom, c)
	excute("Spotify", mt.Spotify, c)
	excute("Steam", mt.Steam, c)
	excute("ChatGPT", mt.ChatGPT, c)
}

func HongKong(c http.Client) {
	// R = append(R, &result{Name: "Hong Kong", Divider: true})
	excute("Now E", mt.NowE, c)
	excute("Viu.TV", mt.ViuTV, c)
	excute("MyTVSuper", mt.MyTvSuper, c)
	excute("HBO GO Aisa", mt.HboGoAisa, c)
	excute("BiliBili Hongkong/Macau Only", mt.BilibiliHKMC, c)
}

func Taiwan(c http.Client) {
	// R = append(R, &result{Name: "Taiwan", Divider: true})
	excute("KKTV", mt.KKTV, c)
	excute("LiTV", mt.LiTV, c)
	excute("MyVideo", mt.MyVideo, c)
	excute("4GTV", mt.TW4GTV, c)
	excute("LineTV", mt.LineTV, c)
	excute("Hami Video", mt.HamiVideo, c)
	excute("CatchPlay+", mt.Catchplay, c)
	excute("Bahamut Anime", mt.BahamutAnime, c)
	excute("HBO GO Aisa", mt.HboGoAisa, c)
	excute("Bilibili Taiwan Only", mt.BilibiliTW, c)
}

func Japan(c http.Client) {
	// R = append(R, &result{Name: "Japan", Divider: true})
	excute("DMM", mt.DMM, c)
	excute("Abema", mt.Abema, c)
	excute("Niconico", mt.Niconico, c)
	excute("music.jp", mt.MusicJP, c)
	excute("Telasa", mt.Telasa, c)
	excute("Paravi", mt.Paravi, c)
	excute("U-NEXT", mt.U_NEXT, c)
	excute("Hulu Japan", mt.HuluJP, c)
	excute("GYAO!", mt.GYAO, c)
	excute("VideoMarket", mt.VideoMarket, c)
	excute("FOD(Fuji TV)", mt.FOD, c)
	excute("Radiko", mt.Radiko, c)
	excute("Karaoke@DAM", mt.Karaoke, c)
	excute("J:COM On Demand", mt.J_COM_ON_DEMAND, c)
	excute("Kancolle", mt.Kancolle, c)
	excute("Pretty Derby Japan", mt.PrettyDerbyJP, c)
	excute("Konosuba Fantastic Days", mt.KonosubaFD, c)
	excute("Princess Connect Re:Dive Japan", mt.PCRJP, c)
	excute("World Flipper Japan", mt.WFJP, c)
	excute("Project Sekai: Colorful Stage", mt.PJSK, c)
}

func NorthAmerica(c http.Client) {
	// R = append(R, &result{Name: "North America", Divider: true})
	excute("FOX", mt.Fox, c)
	excute("Hulu", mt.Hulu, c)
	excute("ESPN+", mt.ESPNPlus, c)
	excute("Epix", mt.Epix, c)
	excute("Starz", mt.Starz, c)
	excute("Philo", mt.Philo, c)
	excute("FXNOW", mt.FXNOW, c)
	excute("TLC GO", mt.TlcGo, c)
	excute("HBO Max", mt.HBOMax, c)
	excute("Shudder", mt.Shudder, c)
	excute("BritBox", mt.BritBox, c)
	excute("CW TV", mt.CW_TV, c)
	excute("NBA TV", mt.NBA_TV, c)
	excute("Fubo TV", mt.FuboTV, c)
	excute("Tubi TV", mt.TubiTV, c)
	excute("Sling TV", mt.SlingTV, c)
	excute("Pluto TV", mt.PlutoTV, c)
	excute("Acorn TV", mt.AcornTV, c)
	excute("SHOWTIME", mt.SHOWTIME, c)
	excute("encoreTVB", mt.EncoreTVB, c)
	excute("Funimation", mt.Funimation, c)
	excute("Discovery+", mt.DiscoveryPlus, c)
	excute("Paramount+", mt.ParamountPlus, c)
	excute("Peacock TV", mt.PeacockTV, c)
	excute("Popcornflix", mt.Popcornflix, c)
	excute("Crunchyroll", mt.Crunchyroll, c)
	excute("Direct Stream", mt.DirectvStream, c)
	// R = append(R, &result{Name: "CA", Divider: true})
	excute("CBC Gem", mt.CBCGem, c)
	excute("Crave", mt.Crave, c)
}

func Check() {

}
