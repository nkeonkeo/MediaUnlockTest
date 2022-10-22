#!/bin/sh
if [ -f "/usr/bin/unlock-test" ];then
    unlock-test -u
    unlock-test $*
    exit
fi

os=`uname -s | tr [:upper:] [:lower:]`
arch=`uname -m`

case ${arch} in
x86)
arch="386"
;;
x86_64)
arch="amd64"
;;
aarch64)
arch="arm64"
;;
esac
# url="https://github.com/nkeonkeo/MediaUnlockTest/releases/latest/download/unlock-test_"${os}"_"${arch}
url="https://unlock.moe/latest/unlock-test_"${os}"_"${arch}
wget ${url} -O unlock-test || curl ${url} -o unlock-test
chmod +x unlock-test
mv unlock-test /usr/bin/unlock-test
unlock-test -v && echo "unlock-test 安装成功" && unlock-test $*
