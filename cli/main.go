package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	m "MediaUnlockTest"

	pb "github.com/schollz/progressbar/v3"
)

var IPV4 = true
var IPV6 = true
var M, TW, HK, JP bool

type result struct {
	Name    string
	Divider bool
	Value   m.Result
}

var tot int64
var R []*result
var bar *pb.ProgressBar
var wg *sync.WaitGroup

func excute(Name string, F func(client http.Client) m.Result, client http.Client) {
	r := &result{Name: Name}
	R = append(R, r)
	wg.Add(1)
	go func() {
		r.Value = F(client)
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
			s += " (region: " + r.Region + ")"
		}
		s += FontSuffix
	} else {
		if r.Err != nil {
			return FontYellow + "ERR: (" + r.Err.Error() + ")" + FontSuffix
		} else if r.Info != "" {
			return FontRed + "NO" + FontSuffix + FontYellow + " (" + r.Info + ")" + FontSuffix
		} else {
			return FontRed + "NO" + FontSuffix
		}
	}
	return
}

func ShowR() {
	NameLength := 25
	for _, r := range R {
		if len(r.Name) > NameLength {
			NameLength = len(r.Name)
		}
	}
	for _, r := range R {
		if r.Divider {
			s := "[ " + r.Name + " ] "
			for i := NameLength - len(s) + 4; i > 0; i-- {
				s += "="
			}
			fmt.Println(s)
		} else {
			result := ShowResult(r.Value)
			if r.Value.Success && strings.HasSuffix(r.Name, "CDN") {
				result = FontSkyBlue + r.Value.Region + FontSuffix
			}
			fmt.Printf("%-"+strconv.Itoa(NameLength)+"s %s\n", r.Name, result)
		}
	}
}

func NewBar(count int64) *pb.ProgressBar {
	return pb.NewOptions64(
		count,
		pb.OptionSetDescription("testing"),
		pb.OptionSetWriter(os.Stderr),
		pb.OptionSetWidth(20),
		pb.OptionThrottle(100*time.Millisecond),
		pb.OptionShowCount(),
		pb.OptionClearOnFinish(),
		pb.OptionEnableColorCodes(true),
		// pb.OptionOnCompletion(func() { bar.Clear() }),
		pb.OptionSpinnerType(14),
		// pb.OptionSetRenderBlankState(true),
	)
}

func Multination(c http.Client) {
	R = append(R, &result{Name: "Multination", Divider: true})
	excute("Dazn", m.Dazn, c)
	excute("Hotstar", m.Hotstar, c)
	excute("Disney+", m.DisneyPlus, c)
	excute("Netflix", m.NetflixRegion, c)
	excute("Netflix CDN", m.NetflixCDN, c)
	excute("Youtube", m.YoutubeRegion, c)
	excute("Youtube CDN", m.YoutubeCDN, c)
	excute("Amazon Prime Video", m.PrimeVideo, c)
	excute("TVBAnywhere+", m.TVBAnywhere, c)
	excute("iQyi", m.IqRegion, c)
	excute("Viu.com", m.ViuCom, c)
	excute("Spotify", m.Spotify, c)
	excute("Steam", m.Steam, c)
}

func HongKong(c http.Client) {
	R = append(R, &result{Name: "Hong Kong", Divider: true})
	excute("Now E", m.NowE, c)
	excute("Viu.TV", m.ViuTV, c)
	excute("MyTVSuper", m.MyTvSuper, c)
	excute("HBO GO Aisa", m.HboGoAisa, c)
	excute("BiliBili Hongkong/Macau/Taiwan", m.BilibiliHKMCTW, c)
}

func Taiwan(c http.Client) {
	R = append(R, &result{Name: "Taiwan", Divider: true})
	excute("KKTV", m.KKTV, c)
	excute("LiTV", m.LiTV, c)
	excute("MyVideo", m.MyVideo, c)
	excute("4GTV", m.TW4GTV, c)
	excute("LineTV", m.LineTV, c)
	excute("Hami Video", m.HamiVideo, c)
	excute("CatchPlay+", m.Catchplay, c)
	excute("Bahamu Anime", m.BahamuAnime, c)
	excute("HBO GO Aisa", m.HboGoAisa, c)
	excute("Bilibili Taiwan Only", m.BilibiliTW, c)
}

func Japan(c http.Client) {
	R = append(R, &result{Name: "Japan", Divider: true})
	excute("DMM", m.DMM, c)
	excute("Abema", m.Abema, c)
	excute("Niconico", m.Niconico, c)
	excute("music.jp", m.MusicJP, c)
	excute("Telasa", m.Telasa, c)
	excute("Paravi", m.Paravi, c)
	excute("U-NEXT", m.U_NEXT, c)
	excute("Hulu Japan", m.HuluJP, c)
	excute("GYAO!", m.GYAO, c)
	excute("VideoMarket", m.VideoMarket, c)
	excute("FOD(Fuji TV)", m.FOD, c)
	excute("Radiko", m.Radiko, c)
	excute("Karaoke@DAM", m.Karaoke, c)
	excute("J:COM On Demand", m.J_COM_ON_DEMAND, c)
	excute("Kancolle", m.Kancolle, c)
	excute("Pretty Derby Japan", m.PrettyDerbyJP, c)
	excute("Konosuba Fantastic Days", m.KonosubaFD, c)
	excute("Princess Connect Re:Dive Japan", m.PCRJP, c)
	excute("World Flipper Japan", m.WFJP, c)
	excute("Project Sekai: Colorful Stage", m.PJSK, c)
}

func Ipv6Multination() {
	c := m.Ipv6HttpClient
	R = append(R, &result{Name: "IPV6 Multination", Divider: true})
	excute("Hotstar", m.Hotstar, c)
	excute("Disney+", m.DisneyPlus, c)
	excute("Netflix", m.NetflixRegion, c)
	excute("Netflix CDN", m.NetflixCDN, c)
	excute("Youtube", m.YoutubeRegion, c)
	excute("Youtube CDN", m.YoutubeCDN, c)
}

func GetIpInfo() {
	resp, err := m.Ipv4HttpClient.Get("https://www.cloudflare.com/cdn-cgi/trace")
	if err != nil {
		IPV4 = false
		fmt.Println("unsupport ipv4")
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		IPV4 = false
		fmt.Println("unsupport ipv4")
	}
	s := string(b)
	i := strings.Index(s, "ip=")
	s = s[i+3:]
	i = strings.Index(s, "\n")
	fmt.Println("Your IPV4 address:", FontSkyBlue, s[:i], FontSuffix)
	resp, err = m.Ipv6HttpClient.Get("https://www.cloudflare.com/cdn-cgi/trace")
	if err != nil {
		IPV6 = false
		fmt.Println("unsupport ipv6")
		return
	}
	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("unsupport ipv6")
	}
	s = string(b)
	i = strings.Index(s, "ip=")
	s = s[i+3:]
	i = strings.Index(s, "\n")
	fmt.Println("Your IPV6 address:", FontSkyBlue, s[:i], FontSuffix)
}

func ReadSelect() {
	fmt.Println("请选择检测项目,直接按回车将进行全部检测: ")
	fmt.Println("[0]: 跨国平台")
	fmt.Println("[1]: 台湾平台")
	fmt.Println("[2]: 香港平台")
	fmt.Println("[3]: 日本平台")
	fmt.Print("请输入对应数字,空格分隔(回车确认): ")
	r := bufio.NewReader(os.Stdin)
	l, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	for _, c := range strings.Split(string(l), " ") {
		switch c {
		case "0":
			M = true
		case "1":
			TW = true
		case "2":
			HK = true
		case "3":
			JP = true
		default:
			M, TW, HK, JP = true, true, true, true
		}
	}
}

func main() {
	client := m.AutoHttpClient
	mode := 0
	flag.IntVar(&mode, "m", 0, "mode 0(default)/4/6")
	flag.Parse()
	if mode == 4 {
		client = m.Ipv4HttpClient
		IPV6 = false
	}
	if mode == 6 {
		client = m.Ipv6HttpClient
		IPV4 = false
		M = true
	}

	GetIpInfo()
	if IPV4 {
		ReadSelect()
	}

	if IPV4 && M {
		tot += 13
	}
	if IPV4 && TW {
		tot += 10
	}
	if IPV4 && HK {
		tot += 5
	}
	if IPV4 && JP {
		tot += 20
	}
	if IPV6 && M {
		tot += 6
	}
	wg = &sync.WaitGroup{}
	bar = NewBar(tot)
	if IPV4 && M {
		Multination(client)
	}
	if IPV4 && TW {
		Taiwan(client)
	}
	if IPV4 && HK {
		HongKong(client)
	}
	if IPV4 && JP {
		Japan(client)
	}
	if IPV6 && M {
		Ipv6Multination()
	}

	wg.Wait()
	bar.Finish()
	ShowR()
}
