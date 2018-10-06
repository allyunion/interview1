package main

import (
  "flag"
  "fmt"
  "net/http"
  "time"
)

func webTime(w http.ResponseWriter, r *http.Request) {
  t := time.Now().UTC()
  message := fmt.Sprintf("{ \"current_time\": %s }", t.Format(time.RFC1123))
  w.Write([]byte(message))
}

func main() {
  portPtr := flag.Int("port", 8080, "port to use")
  flag.Parse()
  port := fmt.Sprintf(":%v", *portPtr)
  fmt.Printf("Starting service on port %v...\n", *portPtr)
  http.HandleFunc("/", webTime)
  if err := http.ListenAndServe(port, nil); err != nil {
    panic(err)
  }
}
