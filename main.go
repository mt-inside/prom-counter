package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	// TODO: use a vec counter and use a label for each http path (this is probably a library function)
	rootHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "root_hits",
			Help: "Requests to /",
		},
	)
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("hit root")
	rootHits.Inc()
	fmt.Fprintf(w, "hi")
}

func init() {
	prometheus.MustRegister(rootHits)
}

func main() {
	log.Println("vim-go")
	http.HandleFunc("/", rootHandler)
	// TODO: move onto different port if that's the convention
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
