package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "flag"
    "strings"
    "math"
    "time"
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
    symbol := flag.String("s","USD","refer to : https://blockchain.info/ticker")
    count := flag.Float64("c",0.0,"count of btcoin")
    silent := flag.Bool("m",false,"silent model ,without loading output")
    flag.Parse()
    symbols := strings.Split(*symbol, ",")
    go displayLoading(*silent)
    resp, err := http.Get("https://blockchain.info/ticker")
    if err != nil {}
    if resp == nil {}
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {}
    var result map[string]interface{}
    err1 := json.Unmarshal(body,&result)
    if err1 != nil {return}

    for key,value := range result{
        for _,value2 := range symbols {
            if key != string(value2) {continue}
            ws := value.(interface{})
            wsMap := ws.(map[string]interface{})
            if vCw, ok := wsMap["15m"]; ok {
                cw := vCw.(float64)
                if math.Dim(*count, 0.0) <= 0.0 {
                    fmt.Printf("\r%v %.2f     \n", wsMap["symbol"] , cw)
                }else{
                    fmt.Printf("\r%v %.2f %.2f     \n ", wsMap["symbol"] , cw , *count*cw)
                }
            }
        }
    }


}