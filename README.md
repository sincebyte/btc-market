# btc-market
go语言获取比特币行情，bit coin market ,比特币行情

# 使用
-s 换算单位，可以设置多个<br>
-c 比特币数量
```
./btcprice:
  -c float
    	count of btcoin
  -m silent model ,without loading output
  -s string
    	refer to : https://blockchain.info/ticker (default "USD")
```
# 返回
换算单位标识符 , 单价 , 比特币数量价值

```
root@mac > ./btcprice                           
$ 7413.92

root@mac > ./btcprice -s CNY                               
¥ 46429.95

root@mac > ./btcprice -s CNY,USD -c 0.044454                               
¥ 46429.95 2064.00
$ 7387.46 328.40
```
# 编译
go build ./btcprice.go

## 跨平台交叉编译
Linux<br>
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

## 跨平台交叉编译,有效组合：
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

# 数据来源
https://blockchain.info/ticker
```
{
  "USD" : {"15m" : 7392.1196853, "last" : 7392.1196853, "buy" : 7392.1196853, "sell" : 7392.1196853, "symbol" : "$"},
  "AUD" : {"15m" : 9614.19086270118, "last" : 9614.19086270118, "buy" : 9614.19086270118, "sell" : 9614.19086270118, "symbol" : "$"},
  "BRL" : {"15m" : 24473.82985409124, "last" : 24473.82985409124, "buy" : 24473.82985409124, "sell" : 24473.82985409124, "symbol" : "R$"},
  "CAD" : {"15m" : 9513.680211340155, "last" : 9513.680211340155, "buy" : 9513.680211340155, "sell" : 9513.680211340155, "symbol" : "$"},
  "CHF" : {"15m" : 7050.832911549384, "last" : 7050.832911549384, "buy" : 7050.832911549384, "sell" : 7050.832911549384, "symbol" : "CHF"},
  "CLP" : {"15m" : 4476371.996630267, "last" : 4476371.996630267, "buy" : 4476371.996630267, "sell" : 4476371.996630267, "symbol" : "$"},
  "CNY" : {"15m" : 46459.22828216088, "last" : 46459.22828216088, "buy" : 46459.22828216088, "sell" : 46459.22828216088, "symbol" : "¥"},
  "DKK" : {"15m" : 44703.10458488322, "last" : 44703.10458488322, "buy" : 44703.10458488322, "sell" : 44703.10458488322, "symbol" : "kr"},
  "EUR" : {"15m" : 6014.0386327, "last" : 6014.0386327, "buy" : 6014.0386327, "sell" : 6014.0386327, "symbol" : "€"},
  "GBP" : {"15m" : 5249.868616260689, "last" : 5249.868616260689, "buy" : 5249.868616260689, "sell" : 5249.868616260689, "symbol" : "£"},
  "HKD" : {"15m" : 58021.76013031659, "last" : 58021.76013031659, "buy" : 58021.76013031659, "sell" : 58021.76013031659, "symbol" : "$"},
  "INR" : {"15m" : 480543.2204421397, "last" : 480543.2204421397, "buy" : 480543.2204421397, "sell" : 480543.2204421397, "symbol" : "₹"},
  "ISK" : {"15m" : 729232.606954845, "last" : 729232.606954845, "buy" : 729232.606954845, "sell" : 729232.606954845, "symbol" : "kr"},
  "JPY" : {"15m" : 789070.3094152, "last" : 789070.3094152, "buy" : 789070.3094152, "sell" : 789070.3094152, "symbol" : "¥"},
  "KRW" : {"15m" : 7803786.830574358, "last" : 7803786.830574358, "buy" : 7803786.830574358, "sell" : 7803786.830574358, "symbol" : "₩"},
  "NZD" : {"15m" : 10196.342464277612, "last" : 10196.342464277612, "buy" : 10196.342464277612, "sell" : 10196.342464277612, "symbol" : "$"},
  "PLN" : {"15m" : 25264.417054434074, "last" : 25264.417054434074, "buy" : 25264.417054434074, "sell" : 25264.417054434074, "symbol" : "zł"},
  "RUB" : {"15m" : 425179.94005908543, "last" : 425179.94005908543, "buy" : 425179.94005908543, "sell" : 425179.94005908543, "symbol" : "RUB"},
  "SEK" : {"15m" : 61881.119288934555, "last" : 61881.119288934555, "buy" : 61881.119288934555, "sell" : 61881.119288934555, "symbol" : "kr"},
  "SGD" : {"15m" : 9681.902679018529, "last" : 9681.902679018529, "buy" : 9681.902679018529, "sell" : 9681.902679018529, "symbol" : "$"},
  "THB" : {"15m" : 230560.212984507, "last" : 230560.212984507, "buy" : 230560.212984507, "sell" : 230560.212984507, "symbol" : "฿"},
  "TWD" : {"15m" : 215481.30893901156, "last" : 215481.30893901156, "buy" : 215481.30893901156, "sell" : 215481.30893901156, "symbol" : "NT$"}
}
```

