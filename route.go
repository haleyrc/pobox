package pobox

import "net/http"

type Route struct {
	h    http.HandlerFunc
	mw   pipeline
	path path
}

func (rt *Route) build() http.HandlerFunc {
	return rt.mw.Apply(rt.h)
}

// func (rt *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	h := rt.mw.Apply(rt.h.ServeHTTP)
// 	h(w, r)
// }

func (rt *Route) Use(mw ...MiddlewareFunc) {
	rt.mw.Use(mw...)
}
