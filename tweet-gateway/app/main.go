// File: api-gateway/main.go
// Este es un API Gateway simple que actúa como un reverse proxy.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// target representa un microservicio de backend.
type target struct {
	name  string
	url   *url.URL
	proxy *httputil.ReverseProxy
}

// newTarget crea una nueva instancia de un servicio de backend.
func newTarget(name, targetURL string) *target {
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		log.Fatalf("No se pudo parsear la URL para el servicio %s: %v", name, err)
	}
	return &target{
		name:  name,
		url:   parsedURL,
		proxy: httputil.NewSingleHostReverseProxy(parsedURL),
	}
}

type routeMatcher func(r *http.Request) bool

// route asocia una función matcher con un servicio de destino.
type route struct {
	matcher routeMatcher
	target  *target
}

// createRouter crea el handler principal que enruta las peticiones.
func createRouter(routes []route, defaultTarget *target) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, route := range routes {
			if route.matcher(r) {
				log.Printf("Gateway: redirigiendo %s a %s", r.URL.Path, route.target.name)
				route.target.proxy.ServeHTTP(w, r)
				return
			}
		}
		log.Printf("Gateway: redirigiendo %s a %s (default)", r.URL.Path, defaultTarget.name)
		defaultTarget.proxy.ServeHTTP(w, r)
	}
}

// healthHandler es un handler simple para el endpoint de health check.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	userService := newTarget("user-service", "http://user-service:8080")
	followService := newTarget("follow-service", "http://follow-service:8080")
	postService := newTarget("post-service", "http://post-service:8080")
	timelineService := newTarget("timeline-service", "http://timeline-service:8080")

	routes := []route{
		{
			matcher: func(r *http.Request) bool {
				return strings.Contains(r.URL.Path, "/timeline")
			},
			target: timelineService,
		},
		{
			matcher: func(r *http.Request) bool {
				return strings.Contains(r.URL.Path, "/follow") || strings.Contains(r.URL.Path, "/following")
			},
			target: followService,
		},
		{
			matcher: func(r *http.Request) bool {
				return strings.HasPrefix(r.URL.Path, "/api/v1/posts")
			},
			target: postService,
		},
	}

	proxyRouter := createRouter(routes, userService)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.Handle("/", proxyRouter)

	log.Println("API Gateway escuchando en el puerto :8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatalf("No se pudo iniciar el API Gateway: %v", err)
	}
}
