package rest

import (
  "net/http"
)

func secureHeaders(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("X-Xss-Protection", "1; mode=block")
    w.Header().Set("X-Frame-Options", "deny")

    next.ServeHTTP(w, r)
  })
}

func corsHeaders(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")

    if r.Method == http.MethodOptions {
        return
    }

    next.ServeHTTP(w, r)
  })
  
}
