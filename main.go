package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		w.Write([]byte("I don't know the hostname"))
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello from %s", hostname)))
}

func main() {
  port, ok := os.LookupEnv("PORT")
  if !ok {
    port = "8080"
  }

  http.HandleFunc("/hello", helloHandler)
  http.ListenAndServe(":" + port, nil)
}
