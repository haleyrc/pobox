package pobox

import (
	"net/http"

	"github.com/apex/log"
)

type pipeline struct {
	mw []MiddlewareFunc
}

func (p *pipeline) Apply(h http.HandlerFunc) http.HandlerFunc {
	log.WithFields(log.Fields{
		"count": len(p.mw),
	}).Debug("applying middleware")
	for i := len(p.mw); i > 0; i-- {
		h = p.mw[i-1](h)
	}
	return h
}

func (p *pipeline) Use(mw ...MiddlewareFunc) {
	if p.mw == nil {
		p.mw = mw
		return
	}
	p.mw = append(p.mw, mw...)
}
