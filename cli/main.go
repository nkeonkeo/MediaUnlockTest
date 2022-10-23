package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
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
	fmt.Println("测试时间: ", FontYellow+time.Now().Format("2006-01-02 15:04:05")+FontSuffix)
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
		pb.OptionSpinnerType(14),
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
	excute("BiliBili Hongkong/Macau Only", m.BilibiliHKMC, c)
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
	excute("Bahamut Anime", m.BahamutAnime, c)
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
		fmt.Println("No IPv4 support")
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		IPV4 = false
		fmt.Println("No IPv4 support")
	}
	s := string(b)
	i := strings.Index(s, "ip=")
	s = s[i+3:]
	i = strings.Index(s, "\n")
	fmt.Println("Your IPV4 address:", FontSkyBlue, s[:i], FontSuffix)
	resp, err = m.Ipv6HttpClient.Get("https://www.cloudflare.com/cdn-cgi/trace")
	if err != nil {
		IPV6 = false
		fmt.Println("No IPv6 support")
		return
	}
	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("No IPv6 support")
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
		M, TW, HK, JP = true, true, true, true
		return
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

type Downloader struct {
	io.Reader
	Total   uint64
	Current uint64
	Pb      *pb.ProgressBar
	done    bool
}

func (d *Downloader) Read(p []byte) (n int, err error) {
	n, err = d.Reader.Read(p)
	d.Current += uint64(n)
	if d.done {
		return
	}
	d.Pb.Add(n)
	// fmt.Printf("\r正在下载，进度：%.2f%% [%s/%s]", float64(d.Current*10000/d.Total)/100, humanize.Bytes(d.Current), humanize.Bytes(d.Total))
	if d.Current == d.Total {
		d.done = true
		d.Pb.Describe("unlock-test下载完成")
		d.Pb.Finish()
	}
	return
}

func checkUpdate() {
	resp, err := http.Get("https://unlock.moe/latest/version")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	version := string(b)
	if version == m.Version {
		fmt.Println("已经是最新版本")
		return
	}
	fmt.Println("检测到新版本", version)
	OS, ARCH := runtime.GOOS, runtime.GOARCH
	fmt.Println("OS:", OS)
	fmt.Println("ARCH:", ARCH)
	out, err := os.Create("/usr/bin/unlock-test_new")
	if err != nil {
		log.Fatal("[ERR] 创建文件出错:", err)
		return
	}
	defer out.Close()
	log.Println("下载unlock-test中 ...")
	url := "https://unlock.moe/latest/unlock-test_" + runtime.GOOS + "_" + runtime.GOARCH
	resp, err = http.Get(url)
	if err != nil {
		log.Fatal("[ERR] 下载unlock-test时出错:", err)
	}
	defer resp.Body.Close()
	downloader := &Downloader{
		Reader: resp.Body,
		Total:  uint64(resp.ContentLength),
		Pb:     pb.DefaultBytes(resp.ContentLength, "下载进度"),
	}
	if _, err := io.Copy(out, downloader); err != nil {
		log.Fatal("[ERR] 下载unlock-test时出错:", err)
	}
	if os.Chmod("/usr/bin/unlock-test_new", 0777) != nil {
		log.Fatal("[ERR] 更改unlock-test后端权限出错:", err)
	}
	if _, err := os.Stat("/usr/bin/unlock-test"); err == nil {
		if os.Remove("/usr/bin/unlock-test") != nil {
			log.Fatal("[ERR] 删除unlock-test旧版本时出错:", err.Error())
		}
	}
	if os.Rename("/usr/bin/unlock-test_new", "/usr/bin/unlock-test") != nil {
		log.Fatal("[ERR] 更新unlock-test后端时出错:", err)
	}
	log.Println("[OK] unlock-test后端更新成功")
}

func showCounts() {
	resp, err := http.Get("https://unlock.moe/count.php")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	s := strings.Split(string(b), " ")
	d, m, t := s[0], s[1], s[3]
	fmt.Printf("当天运行共%s次, 本月运行共%s次, 共计运行%s次\n", FontSkyBlue+d+FontSuffix, FontYellow+m+FontSuffix, FontGreen+t+FontSuffix)
}

func showAd() {
	resp, err := http.Get("https://unlock.moe/ad")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(b))
}

func main() {
	// m.DisneyPlus(m.AutoHttpClient)
	// return
	client := m.AutoHttpClient
	mode := 0
	showVersion := false
	update := false
	flag.IntVar(&mode, "m", 0, "mode 0(default)/4/6")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&update, "u", false, "update")
	flag.Parse()
	if showVersion {
		fmt.Println(m.Version)
		return
	}
	if update {
		checkUpdate()
		return
	}
	if mode == 4 {
		client = m.Ipv4HttpClient
		IPV6 = false
	}
	if mode == 6 {
		client = m.Ipv6HttpClient
		IPV4 = false
		M = true
	}

	fmt.Println("项目地址: " + FontSkyBlue + "https://github.com/nkeonkeo/MediaUnlockTest" + FontSuffix)
	fmt.Println("使用方式: " + FontYellow + "curl -Ls unlock.moe | sh" + FontSuffix)
	fmt.Println()

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
	fmt.Println()
	ShowR()
	fmt.Println()
	fmt.Println("检测完毕，感谢您的使用!")
	showCounts()
	fmt.Println()
	showAd()
}
