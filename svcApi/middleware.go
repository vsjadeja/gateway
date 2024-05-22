package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/time/rate"
)

const SECRETKEY = `your-256-bit-secret`

var skippedFromAuthRoutes = map[string]string{
	`/user/login`:    `login`,
	`/user/register`: `register`,
	`/user/forgot`:   `forgot`,
	`/metrics`:       `metrics`,
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Bypass open routes from auth
		if _, ok := skippedFromAuthRoutes[r.URL.Path]; ok {
			log.Println(`open route :` + r.URL.Path)
			next.ServeHTTP(w, r)
		} else {
			authHeader := strings.Split(r.Header.Get(`Authorization`), `Bearer `)
			if len(authHeader) != 2 {
				fmt.Println("Malformed token")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Malformed Token"))
				return
			}
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(SECRETKEY), nil
			})
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}
		}
	})
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + `:` + r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func LimiterMiddleware(next http.Handler) http.Handler {
	limit := rate.NewLimiter(1, 2)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limit.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`rate limit exceeded.`))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func EnableCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var origin = ""
		for n, h := range r.Header {
			// get the origin from the request
			if strings.Contains(n, "Origin") {
				for _, h := range h {
					origin = h
				}
			}
		}

		// always allow access origin
		w.Header().Add("Access-Control-Allow-Origin", origin)
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, HEAD, TRACE, DELETE, PATCH, COPY, HEAD, LINK, OPTIONS")

		if r.Method == "OPTIONS" {
			for n, h := range r.Header {
				if strings.Contains(n, "Access-Control-Request") {
					for _, h := range h {
						k := strings.Replace(n, "Request", "Allow", 1)
						w.Header().Add(k, h)
					}
				}
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "gateway_http_duration_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(r.URL.Path))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}
