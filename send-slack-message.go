package main

import(
  "os"
  "io/ioutil"
  "log"
  "net/http"
  "fmt"
  "bytes"
  "encoding/json"
  "time"
)

type page struct {}
func (p page) ServeHTTP (w http.ResponseWriter, r *http.Request){
//   if origin != "https://dimasorokin1987.github.io"{
//     return
//   }
  if origin := r.Header.Get("Origin"); origin != "" {
    w.Header().Set("Access-Control-Allow-Origin", origin)
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers",
        "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  }
  // Stop here if its Preflighted OPTIONS request
  if r.Method == "OPTIONS" {
    return
  }
  
  
  keys, ok := r.URL.Query()["txt"]
  if !ok || len(keys[0]) < 1 {
    log.Fatalln("Url Param 'txt' is missing")
  }
  txt := string(keys[0])
  //fmt.Fprint(w,string(txt))
  //return

  token := os.Getenv("SLACK_SECRET_KEY")
  url := "https://slack.com/api/chat.postMessage"

  //curl -X POST "" -H "accept: application/json" -d token=BOT_ACCESS_TOKEN -d channel=U0G9QF3C6 -d text=Hello -d as_user=true

  requestBody, err := json.Marshal(map[string]string{
  // "token": token,
   "channel": "#general",
   "text": txt,
  })
  if err != nil {
    log.Fatalln(err)
  }

  timeout := time.Duration(5*time.Second)
  client := http.Client{
    Timeout: timeout,
  }
  request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
  if err != nil {
    log.Fatalln(err)
  }
  request.Header.Set("Content-Type","application/json;charset=utf-8")
//Authorization: Bearer xoxp-xxxxxxxxx-xxxx 
  request.Header.Set("Authorization","Bearer "+token)

  //resp, err := http.Get("http://httpbin.org/get")
  //resp, err := http.Post("http://httpbin.org/post","application/json",bytes.NewBuffer(requestBody))
  //resp, err := http.Post(url,"application/json",bytes.NewBuffer(requestBody))
  resp, err := client.Do(request)

  if err != nil {
    log.Fatalln(err)
  }
  defer resp.Body.Close()
  
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }
  //log.Println(string(body))
  //w.Header().Set("Access-Control-Allow-Origin", "*")
  //w.Header().Set("Access-Control-Allow-Origin", "https://dimasorokin1987.github.io")
  fmt.Fprint(w,string(body))
}

func main(){
  var p page
  err := http.ListenAndServe(":"+os.Getenv("PORT"),p)
  if err!=nil {
    log.Fatalln(err)
  }
}
