package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}


func main() {
	port := getEnv("LISTEN_PORT", "8000"))
	fmt.Println("Listening: " + port)

	http.HandleFunc("/", ForwardServer)
	http.ListenAndServe((":" + port), nil)

}

func ForwardServer(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	client := &http.Client{}
	client.Post("https://"+r.URL.Path[1:], "application/json", buf)
	if len(os.Getenv("http_proxy")) == 0 {
		fmt.Println("Forwaring to https://" + r.URL.Path[1:])
	} else {
		fmt.Println("Forwaring to https://" + r.URL.Path[1:] + " through proxy")
	}
}
