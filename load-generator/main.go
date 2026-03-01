package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Contador oficial para o Prometheus
	requestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "loadgen_requests_total",
		Help: "Total de requisições enviadas pelo gerador Go",
	})
)

func init() {
	prometheus.MustRegister(requestsTotal)
}

func main() {
	targetURL := "http://127.0.0.1:8080/ingest"
	payload := []byte(`{"data": "metric-payload", "origin": "go-beast"}`)
	workers := 1000

	// Servidor de métricas para o Prometheus na porta 9090
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		fmt.Println("📊 Servidor de métricas Go ativo em :9090/metrics")
		http.ListenAndServe(":9090", nil)
	}()

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 5000, MaxIdleConnsPerHost: 5000,
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
					requestsTotal.Inc() // Incrementa métrica do Prometheus
				}
			}
		}()
	}
	select {}
}
