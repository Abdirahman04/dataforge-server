package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abdirahman04/dataforge-server/internal/app/handlers"
)

func main() {
  port := ":8989"

  fmt.Println("Server started on port", port)
  
  mux := http.NewServeMux()

  mux.HandleFunc("/",handlers.Info)

  if err := http.ListenAndServe(port, mux);err != nil {
    log.Fatal(err)
  }
}
