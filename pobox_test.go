package pobox_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apex/log"

	"github.com/haleyrc/pobox"
)

func TestSorter(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	s := pobox.NewSorter()
	s.Before(
		func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("s.Before: pre-request")
				next(w, r)
				fmt.Println("s.Before: post-request")
			}
		},
		func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("s.Before: pre-request 2")
				next(w, r)
				fmt.Println("s.Before: post-request 2")
			}
		},
	)
	s.After(
		func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("s.After: pre-request")
				next(w, r)
				fmt.Println("s.After: post-request")
			}
		},
		func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("s.After: pre-request 2")
				next(w, r)
				fmt.Println("s.After: post-request 2")
			}
		},
	)

	s.Post("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Index")
	})

	s.Post("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
	}).Use(
		func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("Route before")
				next(w, r)
				fmt.Println("Route after")
			}
		},
		func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("Route before 2")
				next(w, r)
				fmt.Println("Route after 2")
			}
		},
	)

	srv := httptest.NewServer(s)
	defer srv.Close()

	http.Post(srv.URL, "application/json", nil)
	http.Post(srv.URL+"/", "application/json", nil)
	http.Post(srv.URL+"/hello", "application/json", nil)
}
