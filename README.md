# MediaUnlockTest

更快的流媒体检测工具

## CLI

使用方法: 

```bash
bash <(curl -Ls unlock.moe)
```

只检测IPv4结果:

```bash
bash <(curl -Ls unlock.moe) -m 4
```

只检测IPv6结果：

```bash
bash <(curl -Ls unlock.moe) -m 6
```

|args|description|
|-|-|
|`--dns-servers`|specify dns servers|
|`-I`|bind source ip address / interface|
|`--http-proxy`|set proxy (example: "http://username:password@127.0.0.1:1080")|

## Monitor

请先了解Prometheus和Grafana

[README](https://github.com/nkeonkeo/MediaUnlockTest/blob/main/monitor/readme.md)

## Todo

- 补充对北美、南美、欧洲等地区的解锁检测
- 修复已经存在/可能存在的问题

欢迎提交你的 pull requests

## 二次开发

```golang
import "https://github.com/nkeonkeo/MediaUnlockTest"
```

在你的golang项目中导入即可使用

你可以使用它制作解锁监控等小玩具

## 鸣谢

本项目基于 [lmc的全能检测脚本](https://github.com/lmc999/RegionRestrictionCheck) 的思路使用golang重构，提供更快的检测速度

Powered By [NNC](https://nnc.sh)