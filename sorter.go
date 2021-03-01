package pobox

import "net/http"

func NewSorter() *Sorter {
	return &Sorter{
		before: pipeline{
			mw: make([]MiddlewareFunc, 0),
		},
		after: pipeline{
			mw: make([]MiddlewareFunc, 0),
		},
		routes: make(map[path]*Route),
	}
}

type Sorter struct {
	NotFoundHandler http.HandlerFunc

	before pipeline
	after  pipeline

	routes map[path]*Route
}

func (s *Sorter) Post(p string, h http.HandlerFunc) *Route {
	rt := &Route{
		h:    h,
		path: path(p),
	}
	s.routes[path(p)] = rt
	return rt
}

func (s *Sorter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route, ok := s.routes[pathFromRequest(r)]
	if !ok {
		s.NotFoundHandler(w, r)
		return
	}

	h := s.after.Apply(route.h)
	h = route.mw.Apply(h)
	h = s.before.Apply(h)

	h(w, r)
}

func (s *Sorter) Before(mw ...MiddlewareFunc) {
	s.before.Use(mw...)
}

func (s *Sorter) After(mw ...MiddlewareFunc) {
	s.after.Use(mw...)
}

func pathFromRequest(r *http.Request) path {
	return path(r.URL.Path)
}
