package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
)

func main() {
	appToken := make([]byte, 10)
	_, err := rand.Read(appToken)
	if err != nil {
		panic(err)
	}
	pod := os.Getenv("POD_NAME")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(fmt.Sprintf("Hello! I'm a %s with a token of %s!\n", pod, hex.EncodeToString(appToken))))
		if err != nil {
			fmt.Println("Error!", err)
		}
	})

	err = http.ListenAndServe(":8080", nil)
	panic(err)
}
