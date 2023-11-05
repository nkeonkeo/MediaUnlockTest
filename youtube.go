package mediaunlocktest

import (
	"bufio"
	"io"
	"net/http"
	"strings"
)

func YoutubeRegion(c http.Client) Result {
	resp, err := GET(c, "https://www.youtube.com/red")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "www.google.cn") {
		return Result{Status: StatusNo, Region: "cn"}
	}
	if strings.Contains(s, "Premium is not available in your country") {
		return Result{Status: StatusNo}
	}
	if EndLocation := strings.Index(s, `"countryCode":`); EndLocation != -1 {
		return Result{
			Status: StatusOK,
			Region: strings.ToLower(s[EndLocation+15 : EndLocation+17]),
		}
	}
	if strings.Contains(s, "manageSubscriptionButton") {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusNo}
}

func YoutubeCDN(c http.Client) Result {
	resp, err := GET(c, "https://redirector.googlevideo.com/report_mapping")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	r := bufio.NewReader(resp.Body)
	b, _, err := r.ReadLine()
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	s := string(b)
	i := strings.Index(s, "=> ")
	if i == -1 {
		return Result{Status: StatusUnexpected}
	}
	s = s[i+3:]
	i = strings.Index(s, " ")
	if i == -1 {
		return Result{Status: StatusUnexpected}
	}
	s = s[:i]
	i = strings.Index(s, "-")

	if i == -1 {
		i = strings.Index(s, ".")
		return Result{
			Status: StatusOK,
			Region: findAirCode(s[i+1:]),
			Info:   "Youtube Video Server",
		}
	} else {
		isp := s[:i]
		return Result{
			Status: StatusOK,
			Region: isp + " - " + findAirCode(s[i+1:]),
			Info:   "Google Global CacheCDN (ISP Cooperation)",
		}
	}
}

func findAirCode(code string) string {
	airPortCode := []string{"KIX", "NRT", "GMP", "YOW", "YMQ/YUL", "YVR", "YYC", "YEG", "YTO/YYZ", "WAS/IAD", "ABE", "ABQ", "ATL", "AUS", "AZO", "BDL", "BHM", "BNA", "BOI", "BOS", "BRO", "BTR", "BTL", "BUF", "BWI", "CAE", "CAK", "CHA", "CHI/ORD", "CHS", "CID", "CLE", "CLT", "CMH", "CRP", "CVG", "DAY", "DEN", "DFW", "DSM", "DTW", "ELP", "ERI", "EWR", "EVV", "FLL", "FNT", "FWA", "GRR", "GEG", "GSO", "GSP", "GRB", "HAR", "HOU/IAH", "HSV", "HNL", "ICT", "ILM", "IND", "JAN", "JAX", "LAS", "LAX", "LEX", "LIT", "LNK", "LRD", "MCI", "MCO", "MEM", "MFE", "MIA", "MKC", "MKE", "MSN", "MSP", "MSY", "MOB", "NYC/JFK", "OKC", "OMA", "ORF", "ORL", "PBI", "PDX", "PHL/PHA", "PHX", "PIA", "PIT", "PNS", "PVD", "RDU", "RIC", "RNO", "ROC", "SAN", "SAT", "SAV", "SBN", "SDF", "SEA", "BFI", "SFO", "SGF", "SHV", "SLC", "SMF", "STL", "TUL", "SYR", "TOL", "TPA", "TUL", "TUS", "TYS", "MEX", "GDL/MEX", "GUA", "TGU", "SAL", "MGA", "SJO", "PTY", "NAS", "HAV", "SCU", "KIN", "PAP", "SDQ", "SJU", "ROX", "GND", "BGI", "POS", "BOG", "CCS", "GEO", "PBM", "CAY", "BSB", "CWB", "POA", "MAO", "RIO", "SAO", "UIO", "GYE", "LIM", "SRE", "ASU", "MVD", "BUE", "ANF", "SCL", "PTP", "LON/LHR", "ABZ", "BHX", "BOH", "BRS", "CWL", "EDI", "EXT", "GLA", "LPL", "MAN", "NWI", "PLH", "SOU", "BRS", "CDQ", "CVT", "LBA", "PME", "NCL", "HUY", "PIK", "EMA", "BFS", "DUB", "ORK", "SNN", "BRU", "ANR", "OST", "LUX", "AMS", "RTM", "EIN", "ENS", "CPH", "ALL", "AAR", "BLL", "BER/TXL", "MUC", "BRE", "HAJ", "DUS", "FRA", "LEJ", "DUI", "STR", "HAM", "ERF", "FMO", "NUE", "DRS", "SCN", "CGN", "DTM", "BFE", "ZTZ", "ESS", "BON", "RUN", "PAR/CDG", "MRS", "LYS", "BOD", "LIL", "TLS", "NTE", "MLH", "MPL", "GNB", "URO", "NCE", "SXB", "XVE", "PPT", "XMM/GRZ", "BRN", "GVA", "ZRH", "BSL", "ALV", "MAD", "ALC", "BCN", "VLC", "SVQ", "AGP", "VLL", "LIS", "OPO", "ROM", "AHO", "AOI", "BDS", "BLQ", "BRI", "GOA", "MIL/MXP", "SWK", "NAP", "VCE", "FLR", "TRN", "TRS", "CTA", "TAR", "PSA", "QME", "VRN", "ATH", "SKG", "VIE", "LNZ", "GRZ", "SZG", "INN", "PRG", "HEL", "STO/ARN", "AGH", "GOT", "MMA/MMX", "NRK", "OSL", "TIA", "SKP", "SOF", "BEG", "BUH", "KIV", "ZAG", "LJU", "BUD", "BTS", "WAW", "KRK", "GDN", "VNO", "RIX", "TLL", "REK", "MOW", "LED", "MSQ", "IEV/KBP", "SJJ", "THR", "ABD", "KBL", "KWI", "RUH", "JED", "DMM", "SAH", "ADE", "BGW", "BEY", "BAH", "AUH", "DXB", "SHJ", "DOH", "JRD", "TLV", "DAM", "AMM", "ANK", "ADA", "BTZ", "IZM", "IST", "BAH", "NIC", "LCA", "BAK", "EVN", "TBS", "MSH", "ASB", "DYU", "KGF", "FRU", "TAS", "CAI", "KRT", "MCT", "ADD", "JIB", "NBO", "TIP", "ALG", "AAE", "TUN", "RBA", "CAS", "NDJ", "NIM", "ABV", "LOS", "PHC", "BKO", "OUA", "COO", "LFW", "ACC", "ASK", "ABJ", "HGS", "MLW", "CKF", "DKR", "BJL", "KLA", "BGF", "YAO", "SSG", "KLA", "KGL", "DAR", "BJM", "BZV", "LBV", "TMS", "MPM", "LLW", "LUN", "HRE", "LAD", "GBE", "WDH", "JNB", "DUR", "CPT", "MRU", "TNR", "YVA", "SEZ", "NKC", "HKG", "TPE", "KHH", "FNJ", "SEL/ICN", "PUS", "TYONRT", "KIX/OSA", "NGO", "FUK", "YOK", "HIJ", "OKA", "SDJ", "SPA", "MNL", "HEB", "DVO", "KUL", "PEN", "LGK", "BKI", "KCH", "IPH", "JHB", "KBR", "SBW", "SDK", "BWN", "SIN", "JKT", "MES", "SUB", "DPS", "UPG", "PNK", "DIL", "SGN", "HAN", "HPH", "VTE", "BKK", "CEI", "HDY", "HKT", "NSI", "RGN", "MDL", "PNH", "DAC", "CGP", "DEL", "BOM", "CCU", "MAA", "BLR", "SXM", "HYD", "KTM", "ISB", "KHI", "LHE", "PEW", "CMB", "MLE", "ULN", "CBR", "MEL", "ADL", "DRN", "CNS", "BNE", "PER", "SYD", "WLG", "AKL", "CHC", "POM", "SUV", "TRW", "HIR", "TBU", "APW", "FUN", "KSA", "VLI"}
	// airPortName := []string{"日本 大阪","日本 东京","韩国 金浦", "加拿大 渥太华","加拿大 蒙特利尔","加拿大 温哥华","加拿大 卡尔加里","加拿大 埃德蒙顿","加拿大 多伦多","美国  华盛顿","美国  阿伦敦","美国  阿尔伯克斯","美国  阿特兰大","美国  奥斯汀","美国  卡拉马祖","美国  哈特福德","美国  伯明翰","美国  纳什维尔","美国  博伊西","美国  波士顿","美国  布朗斯韦尔","美国  巴吞鲁日","美国  巴特尔克里克","美国  布法罗","美国  巴尔的摩","美国  哥伦比亚","美国  阿克伦肯顿","美国  查塔诺加","美国  芝加哥","美国  查尔斯顿","美国  锡达拉皮兹","美国  克利夫兰","美国  夏洛特","美国  哥伦布","美国  科珀斯克里斯蒂","美国  辛辛那提","美国  代顿","美国  丹佛","美国  达拉斯","美国  得梅因","美国  底特律","美国  埃尔帕索","美国   伊利/伊利湖","美国  纽瓦克","美国  埃文斯韦尔","美国  劳德代尔堡","美国  弗林特","美国  育空堡","美国  大急流域","美国  斯波坎","美国  格林斯伯勒","美国  格林维尔","美国  格林贝","美国  哈里斯堡","美国  休斯敦","美国  亨茨维尔","美国  火奴鲁鲁（夏威夷州的首府)","美国  威奇托","美国  威尔明顿","美国  印第安纳波利斯","美国  杰克逊","美国  杰克逊威尔","美国  拉斯维加斯","美国  洛杉机","美国  列克星敦","美国  小石城","美国  林肯","美国  拉雷多","美国  堪萨斯城","美国  奥兰多","美国  孟菲斯","美国  麦卡伦","美国  迈阿密","美国  堪萨斯","美国  密尔沃基","美国  麦迪逊","美国  明利阿波利斯 ","美国  新奥尔良","美国  莫比尔","美国  纽约","美国  俄克拉荷马城","美国  奥马哈","美国  诺福克","美国  奥兰多","美国  西棕榈滩","美国  波特兰","美国  费城","美国  费尼克斯","美国  皮奥里亚","美国  匹兹堡","美国  彭萨科拉","美国  普罗维登斯","美国  达拉谟","美国  里士满","美国  里诺","美国  罗彻斯特","美国  圣迭戈","美国  圣安东尼奥","美国  萨凡纳","美国  南本德","美国  路易斯维尔","美国  西雅图","美国  西雅图","美国  旧金山","美国  斯普林菲尔德","美国  什里夫波特","美国  盐湖城","美国  萨克拉门托","美国  圣路易斯","美国  塔尔萨","美国  锡拉丘兹","美国  托莱多","美国  坦帕","美国  塔尔萨","美国  图森","美国  诺科斯韦尔","墨西哥墨西哥城","墨西哥瓜达拉哈拉","危地马拉危地马拉","洪都拉斯  特古西加尔巴","萨尔瓦多圣萨尔瓦多","尼加拉瓜   马拉瓜","哥斯达黎加圣何塞","巴拿马 巴拿马城","巴哈马拿骚","古巴哈瓦那","古巴圣地亚哥","牙买加金斯敦","海地太子港","多米尼加圣多明各","波多黎各 圣胡安","多米尼克罗索","格林纳达圣乔治","巴巴多斯 布里奇顿","特立尼达和多巴哥  西班牙港","哥伦比亚 圣菲波哥达","委内瑞拉加拉加斯","圭亚那乔治敦","苏里南帕拉马里博","法属圭那亚卡宴","巴西 巴西利亚","巴西 库里蒂巴","巴西 阿雷格里港","巴西 马卤斯","巴西 里约热内卢","巴西 圣保罗","厄瓜多尔基多","厄瓜多尔瓜尔基尔","秘鲁 利马","玻利维亚   苏克雷","巴拉圭亚松森","乌拉圭蒙得维的亚","阿根廷布宜诺斯艾利斯","智利 安托法加斯塔","智利 圣地亚哥","拉丁美洲瓜德罗普","英国 伦敦","英国 阿伯丁","英国 伯明翰","英国 伯恩茅斯","英国 布里斯托尔","英国 加地夫","英国 爱丁堡","英国 埃克塞特","英国 格拉斯哥 ","英国 利物浦","英国 曼彻斯特","英国 诺里奇","英国 普利茅斯","英国 南安普敦","英国 布里斯托尔","英国 克里伊登","英国 考文垂","英国 利兹","英国 朴茨茅斯","英国 纽卡斯尔","英国 亨伯赛德郡","英国 格拉斯哥 ","英国 诺丁汉","北爱尔兰 贝尔法斯特","爱尔兰 都柏林 ","爱尔兰 科克","爱尔兰 香农","比利时布鲁塞尔","比利时安特卫普","比利时奥斯坦德","卢森堡卢森堡","荷兰阿姆斯特丹","荷兰鹿特丹","荷兰爱恩德霍芬","荷兰恩斯赫德","丹麦 哥本哈根","丹麦 阿尔伯格","丹麦 奥胡斯","丹麦 比灵顿","德国柏林","德国慕尼黑","德国不来梅","德国汉诺威","德国杜塞尔多夫","德国法兰克福","德国莱比锡","德国杜伊斯堡","德国斯图加特","德国汉堡","德国爱尔福特","德国明斯特","德国纽伦堡","德国德累斯顿","德国萨尔布吕肯","德国科隆","德国多特蒙德","德国比勒费尔德","德国开姆尼茨","德国埃森","德国波恩","德国重聚","法国 巴黎","法国 马塞","法国 里昂","法国 波尔多","法国 里尔","法国 图卢兹","法国 南特","法国 牟罗兹","法国 蒙彼利埃","法国 格勒诺布尔","法国 鲁昂","法国 尼斯","法国 斯特拉斯堡","法国 凡尔塞","法属波利尼帕皮提 ","摩纳哥摩纳哥","瑞士伯尔尼","瑞士日内瓦","瑞士苏黎世","瑞士巴塞尔","安道尔安道尔","西班牙马德里","西班牙阿利坎特","西班牙巴塞罗那","西班牙华伦西亚(巴伦比亚)","西班牙塞尔维亚","西班牙马拉加","西班牙巴利亚多利德","葡萄牙里斯本","葡萄牙波尔图","意大利罗马","意大利安齐奥","意大利安科纳","意大利布林迪西","意大利博洛尼亚","意大利巴里","意大利热那亚","意大利米兰","意大利米兰市郊","意大利拿坡里","意大利威尼斯","意大利佛罗伦萨","意大利都灵","意大利的里雅斯特","意大利卡塔尼亚","意大利塔兰托","意大利比萨","意大利墨西拿","意大利维罗纳","希腊雅典","希腊塞萨洛尼基","奥地利维也纳","奥地利林茨","奥地利格拉茨","奥地利萨尔斯堡","奥地利因斯布鲁克","捷克 布拉格","芬兰赫尔辛基","瑞典斯德哥尔摩","瑞典海尔辛堡","瑞典哥德堡","瑞典马尔默","瑞典北雪平","挪威奥斯陆","阿尔巴尼亚地拉那","马其顿斯科普里","保加利亚索非亚","南斯拉夫 贝尔格莱德","罗马尼亚 布加勒斯特","摩尔多瓦 基希纳乌","克罗地亚 萨格勒布","斯洛文尼亚卢布尔雅那","匈牙利布达佩斯","斯洛伐克布拉迪斯拉发","波兰 华沙","波兰 克拉克夫","波兰 格但斯克","立陶宛维尔纽斯","拉托维亚里加","爱沙尼亚 塔林","冰岛雷克亚未克","俄罗斯 莫斯科","俄罗斯 圣彼得堡(列宁格勒)","白俄罗斯明斯克","乌克兰基辅","波黑萨拉热窝","伊朗 德黑兰","伊朗 阿巴达","阿富汗喀布尔","科威特 科威特","沙特阿拉伯利雅得","沙特阿拉伯吉达","沙特阿拉伯达曼","也门萨那","也门亚丁","伊拉克巴格达","黎巴嫩 贝鲁特","黎巴嫩 巴林","阿联酋阿布扎比","阿联酋迪拜","阿联酋沙加","卡塔尔 多哈","以色列  耶路撒冷","以色列  特拉维夫","叙利亚大马士革","约旦 安曼","土耳其安卡拉","土耳其阿达那","土耳其布尔萨","土耳其伊兹密尔","土耳其伊斯坦布尔","巴林  巴林","塞浦路斯尼科西亚","塞浦路斯拉纳卡","阿塞拜疆 巴库","亚美尼亚 埃里温","格鲁吉亚 第比利斯","阿曼  马斯喀特","土库曼斯坦阿什哈巴德","塔吉克斯坦杜尚别","哈萨克斯坦卡拉干达","吉尔吉斯斯坦比什凯克","乌兹别克斯坦 塔什干","埃及开罗","苏丹 喀土穆","苏丹 马斯喀特阿曼","埃塞俄比亚亚的斯亚贝巴","吉布提 吉布提","肯尼亚 内罗毕","利比亚 的黎波里","阿尔及利亚  阿尔及尔","阿尔及利亚  安纳巴","突尼斯  突尼斯","摩洛哥拉巴特","摩洛哥卡萨布兰卡","乍得 恩贾梅纳","尼日尔 尼亚美","尼日利亚 阿布贾","尼日利亚 拉各斯","尼日利亚 哈科特港 ","马里 巴马科","布基纳法索瓦加杜古","贝宁 科托努","多哥 洛美","加纳  阿克拉","科特迪瓦阿穆苏克罗","科特迪瓦阿比让","塞拉利昂弗里敦","利比里亚蒙罗维亚","几内亚科纳克里","塞内加尔 达喀尔","冈比亚 班珠尔","马里塔尼亚努瓦克肖特","中非共和国班吉","喀麦隆雅温得","赤道几内亚马拉博","乌干达坎帕拉","卢旺达 基加利","坦桑尼亚达累斯萨拉姆","布隆迪布琼布拉","刚果布拉柴维尔","加蓬 利伯维尔","圣多美和普林西比圣多美","莫桑比克马普托","马拉维利隆圭","赞比亚卢萨卡","津巴布韦哈拉雷","安哥拉罗安达","博茨瓦纳哈伯罗内","纳米比亚温得和克","南非约翰内斯堡","南非德班","南非开普敦","毛里求斯 毛里求斯","马达加斯加塔那那利佛","科摩罗莫罗尼","马埃塞舌尔群岛","努瓦克肖特毛里塔尼亚","中国香港","中国台北","中国高雄","朝鲜 平壤","韩国 汉城仁川","韩国 釜山","日本 东京","日本 大阪","日本 名古屋","日本 福冈","日本 横滨","日本 广岛","日本 冲绳岛","日本 仙台","日本 札幌","菲律宾马尼拉","菲律宾宿务","菲律宾达沃","马来西亚吉隆坡","马来西亚槟城","马来西亚凌家卫岛","马来西亚哥大基纳巴卢","马来西亚古晋","马来西亚怡宝","马来西亚新山","马来西亚哥打巴鲁","马来西亚诗巫","马来西亚山打根","文莱 斯里巴加湾市","新加坡 新加坡/樟宜  ","印度尼西亚 雅加达","印度尼西亚 棉兰","印度尼西亚 泗水","印度尼西亚 登巴萨","印度尼西亚 乌戒潘当","印度尼西亚 坤甸","东帝汶帝力","越南胡志明市","越南河内","越南海防","老挝 万象","泰国曼谷","泰国清迈","泰国合艾","泰国普吉","泰国","缅甸 仰光","缅甸 曼德勒","柬埔寨 金边","孟加拉达卡","孟加拉吉大港","印度 新德里","印度 孟买","印度 加尔各答","印度 马德拉斯","印度 班加罗尔","印度 荷兰安得列斯群岛","印度 海德拉巴","尼泊尔加德满都","巴基斯坦伊斯兰堡","巴基斯坦卡拉奇","巴基斯坦拉合尔","巴基斯坦白沙瓦","斯里兰卡 科伦坡","马尔代夫 马累","蒙古乌兰巴托","澳大利亚 堪培拉","澳大利亚 墨尔本","澳大利亚 阿德莱得","澳大利亚 达尔文","澳大利亚 凯恩斯","澳大利亚 布里斯班","澳大利亚 珀斯","澳大利亚 悉尼","新西兰惠灵顿","新西兰奥克兰 ","新西兰利特尔顿(基督城)","巴布亚新几内亚莫尔兹比港","斐济群岛 苏瓦","基里巴斯 塔拉瓦","所罗门群岛霍尼亚拉","汤加 努瓦阿洛法","萨摩亚  阿皮亚","图瓦卢富纳富提","密克罗尼西亚帕利基尔","瓦努阿图维拉港"}
	i, v := 0, ""
	for ; i < len(code); i++ {
		if code[i] >= '0' && code[i] <= '9' {
			break
		}
	}
	code = strings.ToUpper(code[:i])
	for i, v = range airPortCode {
		if strings.Contains(code, v) {
			return v
			// return airPortCode[i]
		}
	}
	return code
}
