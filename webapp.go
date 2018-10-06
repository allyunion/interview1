package main

import (
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
  http.HandleFunc("/", webTime)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
