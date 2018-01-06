package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hits",
			Help: "Requests to path",
		},
		[]string{"path"},
	)
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("hit %s", r.URL.Path)
	hits.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	fmt.Fprintf(w, "hi")
}

func init() {
	prometheus.MustRegister(hits)
}

func main() {
	metrics_mux := http.NewServeMux()
	metrics_mux.Handle("/metrics", promhttp.Handler())
	metrics_server := http.Server{
		Addr:    ":8085",
		Handler: metrics_mux,
	}
	go metrics_server.ListenAndServe()
	log.Println("Serving /metrics on :8085")

	main_mux := http.NewServeMux()
	main_mux.HandleFunc("/", rootHandler)
	main_server := http.Server{
		Addr:    ":8080",
		Handler: main_mux,
	}
	log.Println("Serving on :8080")
	log.Fatal(main_server.ListenAndServe())
}
