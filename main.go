package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("Port:", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
