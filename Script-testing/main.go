package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  //"github.com/ararchch/api-gateway/Script-testing"
)

func main() {

  url := "http://127.0.0.1:8080/multiply"
  method := "POST"
  fmt.Println("here1")
  payload := strings.NewReader(`{` + "" +` "FirstNum": "6",`+ "" +` "SecondNum": "3"`+""+`}`)
  fmt.Println("here2")
  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println("errres 1")
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println("err res")
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println("Error body")
    return
  }
  fmt.Println("here")
  fmt.Println(string(body))
  fmt.Println(res.StatusCode)
}