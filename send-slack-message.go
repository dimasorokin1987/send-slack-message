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
  //key := os.Getenv("SLACK_SECRET_KEY")
  //url := 

  requestBody, err := json.Marshall(map[string]string{
   "name": "test",
   "email": "tets@mail.com",
  })
  //resp, err := http.Get("http://httpbin.org/get")
  resp, err := http.Post("http://httpbin.org/post",
   "application/json",
   bytes.NewBuffer(requestBody))
  
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
