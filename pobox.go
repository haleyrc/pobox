package pobox

import (
	"net/http"
)

type path string

type Middleware func(next http.Handler) http.Handler

type MiddlewareFunc func(next http.HandlerFunc) http.HandlerFunc
