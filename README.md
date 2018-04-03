# btc-market
go语言获取比特币行情，bit coin market ,比特币行情

#使用
```
./btcprice:
  -c float
    	count of btcoin
  -s string
    	refer to : https://blockchain.info/ticker (default "USD")
```

#数据来源
https://blockchain.info/ticker

#编译
跨平台交叉编译
Linux
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

有效组合：
```
    $GOOS       $GOARCH   
    android     arm
    darwin      386
    darwin      amd64
    darwin      arm
    darwin      arm64
    dragonfly   amd64
    freebsd     386
    freebsd     amd64
    freebsd     arm
    linux       386
    linux       amd64
    linux       arm
    linux       arm64
    linux       ppc64
    linux       ppc64le
    linux       mips
    linux       mipsle
    linux       mips64
    linux       mips64le
    netbsd      386
    netbsd      amd64
    netbsd      arm
    openbsd     386
    openbsd     amd64
    openbsd     arm
    plan9       386
    plan9       amd64
    solaris     amd64
    windows     386
    windows     amd64
```
