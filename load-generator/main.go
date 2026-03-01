package main

import (
	"bytes"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
	"fmt"
)

func main() {
	targetURL := "http://127.0.0.1:8080/ingest"
	payload := []byte(`{"data": "max-payload", "origin": "go-beast-mode"}`)
	
	// Aumentando para 1000 workers para maximizar o paralelismo no Termux
	workers := 1000 

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        5000,
			MaxIdleConnsPerHost: 5000,
			DisableCompression:  true,
		},
	}

	var rpsCounter int64
	go func() {
		for {
			time.Sleep(1 * time.Second)
			currentRPS := atomic.SwapInt64(&rpsCounter, 0)
			fmt.Printf("🔥 RPS Atual: %d\n", currentRPS)
		}
	}()

	for i := 0; i < workers; i++ {
		go func() {
			for {
				req, _ := http.NewRequest("POST", targetURL, bytes.NewReader(payload))
				req.Header.Set("Content-Type", "application/json")
				resp, err := client.Do(req)
				if err == nil {
					io.Copy(io.Discard, resp.Body)
					resp.Body.Close()
					atomic.AddInt64(&rpsCounter, 1)
				}
			}
		}()
	}
	select {} 
}
