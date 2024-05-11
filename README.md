# MediaUnlockTest

> 更快的流媒体检测工具

比原版提供更多测试项目

## CLI

使用方法: 

```bash
bash <(curl -Ls unlock.icmp.ing/test.sh)
```

只检测IPv4结果:

```bash
bash <(curl -Ls unlock.icmp.ing/test.sh) -m 4
```

只检测IPv6结果：

```bash
bash <(curl -Ls unlock.icmp.ing/test.sh) -m 6
```

|args|description|
|-|-|
|`--dns-servers`|specify dns servers|
|`-I`|bind source ip address / interface|
|`--http-proxy`|set proxy (example: "http://username:password@127.0.0.1:1080")|

## Monitor

使用 Prometheus 和 Grafana 搭建流媒体解锁监控，效果： [ICMPing](https://icmp.ing/service)。

~~图文教程有空再写，暂时鸽了~~

[README](https://github.com/HsukqiLee/MediaUnlockTest/blob/main/monitor/readme.md)

## Todo

- 补充对北美、南美、欧洲等地区的解锁检测
- 修复已经存在/可能存在的问题

欢迎提交你的 Pull Requests

## 二次开发

```golang
import "https://github.com/HsukqiLee/MediaUnlockTest"
```

在你的golang项目中导入即可使用

你可以使用它制作解锁监控等小玩具

## 参与开发的小伙伴

<!--GAMFC_DELIMITER--><a href="https://github.com/nkeonkeo" title="neko"><img src="https://avatars.githubusercontent.com/u/36293036?v=4" width="50;" alt="neko"/></a>
<a href="https://github.com/HsukqiLee" title="HsukqiLee"><img src="https://avatars.githubusercontent.com/u/79034142?v=4" width="50;" alt="HsukqiLee"/></a>
<a href="https://github.com/xkww3n" title="xkww3n"><img src="https://avatars.githubusercontent.com/u/30206355?v=4" width="50;" alt="xkww3n"/></a>
<a href="https://github.com/oif" title="Neo Zhuo"><img src="https://avatars.githubusercontent.com/u/6374269?v=4" width="50;" alt="Neo Zhuo"/></a>
<a href="https://github.com/edgist" title="edgist"><img src="https://avatars.githubusercontent.com/u/34343603?v=4" width="50;" alt="edgist"/></a>
<a href="https://github.com/iawb-ray" title="iawb-ray"><img src="https://avatars.githubusercontent.com/u/49180084?v=4" width="50;" alt="iawb-ray"/></a><!--GAMFC_DELIMITER_END-->

## 鸣谢

原项目基于 [lmc的全能检测脚本](https://github.com/lmc999/RegionRestrictionCheck) 的思路使用 Golang 重构，提供更快的检测速度。

本项目基于 [MediaUnlockTest](https://github.com/nkeonkeo/MediaUnlockTest) 二次开发，提供更丰富的测试项目。

Made with ❤️ By **Hsukqi Lee**.
