package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "flag"
    "strings"
    "math"
)

func main() {
    symbol := flag.String("s","USD","refer to : https://blockchain.info/ticker")
    count := flag.Float64("c",0.0,"count of btcoin")
    flag.Parse()
    symbols := strings.Split(*symbol, ",")
    resp, err := http.Get("https://blockchain.info/ticker")
    if err != nil {}
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
                    fmt.Printf("%v %.2f\n", wsMap["symbol"] , cw)
                }else{
                    fmt.Printf("%v %.2f %.2f\n", wsMap["symbol"] , cw , *count*cw)
                }
            }
        }
    }

}