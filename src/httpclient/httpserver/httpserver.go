package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/zephinzer/ezpromhttp"
)

var (
	namespace         = "ysharma"
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "ysharma_go_http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"app", "name", "method", "status"})
	httpLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "ysharma_go_http_latency",
		Help:      "Time taken to execute endpoint.",
	}, []string{"app", "name", "method", "status"})
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println(r.RequestURI)
		requestLogger := log.WithFields(log.Fields{
			"method":    r.Method,
			"authority": r.Host,
			"uri":       r.RequestURI,
			"alpn":      r.Proto,
		})
		requestLogger.Info()
		next.ServeHTTP(w, r)
	})
}

func main() {
	log.SetLevel(log.InfoLevel)
	router := mux.NewRouter().StrictSlash(true)
	// router.Use(loggingMiddleware)
	router.Handle("/metrics", promhttp.Handler())
	apis := router.PathPrefix("/v1").Subrouter()
	apis.HandleFunc("/get", get).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", ezpromhttp.InstrumentHandler(router)))
}
