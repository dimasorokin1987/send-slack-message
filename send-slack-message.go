package main

import(
  "os"
  "io/ioutil"
  "log"
  "net/http"
  "fmt"
  "bytes"
  "encoding/json"
)

type page struct {}
func (p page) ServeHTTP (w http.ResponseWriter, _ *http.Request){
  token := os.Getenv("SLACK_SECRET_KEY")
  url := "https://slack.com/api/chat.postMessage"

  //curl -X POST "" -H "accept: application/json" -d token=BOT_ACCESS_TOKEN -d channel=U0G9QF3C6 -d text=Hello -d as_user=true

  requestBody, err := json.Marshal(map[string]string{
   "token": token,
   "channel": "#general",
   "text": "test hello",
  })
  //resp, err := http.Get("http://httpbin.org/get")
  //resp, err := http.Post("http://httpbin.org/post","application/json",bytes.NewBuffer(requestBody))
  resp, err := http.Post(url,"application/json",bytes.NewBuffer(requestBody))
  

  if err != nil {
    log.Fatalln(err)
  }
  defer resp.Body.Close()
  
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }
  //log.Println(string(body))
  fmt.Fprint(w,string(body))
}

func main(){
  var p page
  err := http.ListenAndServe(":"+os.Getenv("PORT"),p)
  if err!=nil {
    log.Fatalln(err)
  }
}
