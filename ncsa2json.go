package main

import (
  "fmt"
  "strings"
  "encoding/json"
  "bufio"
  "os"
)

func main() {
  type Request struct {
    Client string `json:"client"`
    Url string `json:"url"`
    Status string `json:"status"`
    Size string `json:"size"`
    Time string `json:"time"`
    Method string `json:"method"`
    Ua string `json:"user-agent"`
    Referer string `json:"referer"`
  }
  in := bufio.NewReader(os.Stdin)
  for {
    log, err := in.ReadString('\n')
    if err != nil {
      fmt.Println("EOF")
      os.Exit(1) // if you return error
    }
    var arr = strings.Split(log, " ")
    var client string = arr[0]
    var url string = arr[6]
    var status string = arr[8]
    var size string = arr[9]
    var time string = arr[3] + " " + arr[4]
    var method string = arr[5][1:]
    arr = strings.Split(log, "\"")
    var ua string = arr[5]
    var referer string = arr[3]
    
    r := Request {client, url, status, size, time, method, ua, referer}
    b, _ := json.MarshalIndent(r, "", "  ")
    fmt.Println(string(b))

    //for index,element := range arr {
      //fmt.Println(index, element)
    //}
  }
}
