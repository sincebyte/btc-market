package main

import (
	"fmt"
	"encoding/json"
	"github.com/leizongmin/huobiapi"
	"flag"
    "time"
    "net/http"
    "io/ioutil"
    "strings"
)

func displayLoading(silent bool){
    if silent {return}
    loading := []string{"loading    ","loading .  ","loading .. ","loading ..."}
    for i := 0;;{
        fmt.Printf("\r%c[1;0;32m%s%c[0m",0x1B, loading[i],0x1B) 
        time.Sleep(time.Millisecond * 150) 
        i = i%3+1
    }
}

func main() {
	symbol := flag.String("s","eosbtc","refer to : https://api.huobipro.com/v1/common/symbols")
	silent := flag.Bool("m",true,"silent model ,without loading output")
	flag.Parse()
	go displayLoading(*silent)
	if(strings.HasSuffix(*symbol,"btc") == false){
		fmt.Printf("error : Symbol need suffix with btc.Like eosbtc,htbtc")
		return
	}

	client, err := huobiapi.NewClient("****-****-****-****", "****-****-****-****")
	if err != nil {
		panic(err)
	}
	ret, err := client.Request("GET", "/market/history/kline", huobiapi.ParamsData{
		"symbol": *symbol,
		"period": "15min",
		"size" : "1",
	})
	data, err := ret.Get("data").Array()
	if err != nil {
		panic(err)
	}
	for _, v := range data {
		wsMap := v.(map[string]interface{})
		f, _ := wsMap["close"].(json.Number).Float64()
		va := *symbol
		fmt.Printf("\r%v : ￥%.4f       \n",va[:len(va)-3], f*getbtcny())
	}
}


func getbtcny() float64 {
	cw := 0.0
    symbol := "CNY"
    symbols := strings.Split(symbol, ",")
    resp, err := http.Get("https://blockchain.info/ticker")
    if err != nil {}
    if resp == nil {}
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {}
    var result map[string]interface{}
    err1 := json.Unmarshal(body,&result)
    if err1 != nil {return cw}
    for key,value := range result{
        for _,value2 := range symbols {
            if key != string(value2) {continue}
            ws := value.(interface{})
            wsMap := ws.(map[string]interface{})
            if vCw, ok := wsMap["15m"]; ok {
                cw = vCw.(float64)
                fmt.Printf("\rbtc : ￥%.4f       \n", cw)
                return cw
            }
        }
    }
    return cw
}

