# MediaUnlockTest

更快的流媒体检测脚本

使用方法: 

```
bash <(curl -Ls unlock.moe)
```

只检测IPv4结果:

```
bash <(curl -Ls unlock.moe) -m 4
```

只检测IPv6结果：

```
bash <(curl -Ls unlock.moe) -m 6
```

## 鸣谢

本项目基于 [LCM的全能检测脚本](https://github.com/lmc999/RegionRestrictionCheck) 的思路使用golang重构，提供更快的检测速度

Powered By [Neko Neko Cloud](https://nekoneko.cloud)