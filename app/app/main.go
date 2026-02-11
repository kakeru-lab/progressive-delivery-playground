package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	version = getenv("APP_VERSION", "v1")
)

func getenv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "progressive-delivery-playground %s\n", version)
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	// /error?rate=0.2 で 20% くらい 500 を返す（カナリア失敗の再現用）
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		rateStr := r.URL.Query().Get("rate")
		if rateStr == "" {
			rateStr = "0"
		}
		rate, _ := strconv.ParseFloat(rateStr, 64)
		if rand.Float64() < rate {
			http.Error(w, "intentional error\n", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "ok (rate=%s) version=%s\n", rateStr, version)
	})

	port := getenv("PORT", "8080")
	addr := ":" + port
	fmt.Println("listening on", addr, "version=", version)
	_ = http.ListenAndServe(addr, mux)
}
