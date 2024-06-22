package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/apricote/cicd-observability/pkg"
)

func main() {
	log.Println(pkg.HelloWorld())
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
