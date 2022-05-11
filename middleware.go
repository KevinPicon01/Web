package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Checking Log...")
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Checking auth...")
			flag := true
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}
