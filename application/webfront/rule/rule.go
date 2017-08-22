package rule

import (
	"net/http"
	"net/http/httputil"
	"strings"
)

// Rule represents a rule in a configuration file.
type Rule struct {
	Host    string // to match against request Host header
	Forward string // non-empty if reverse proxy
	Serve   string // non-empty if file server
}

func (r *Rule) Match(req *http.Request) bool {
	return req.Host == r.Host || strings.HasSuffix(req.Host, "."+r.Host)
}

func (r *Rule) Handler() http.Handler {
	if h := r.Forward; h != "" {
		// func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request)
		// So here a pointer is returned since it's the pointer who implements http.Handler interface.
		// For the client it does not matter it's a pointer or not, both can be used to call ServeHTTP function. (?)
		return &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				req.URL.Scheme = "http"
				req.URL.Host = h
			},
		}
	}
	if d := r.Serve; d != "" {
		return http.FileServer(http.Dir(d))
	}
	// zero value for an interface (?)
	return nil
}
