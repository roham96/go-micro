package web

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/roham96/go-micro/01_user-qry/service"
)

type Router struct {
	srv    *service.Service
	routes []route
}

func NewRouter(srv *service.Service) *Router {
	r := &Router{srv: service.New()}
	r.makeRoutes()
	return r
}

func (r *Router) makeRoutes() {
	r.routes = []route{
		newRoute("GET", "/api/widgets", apiGetWidgets(r.srv)),
		newRoute("POST", "/api/widgets", apiCreateWidget(r.srv)),
		newRoute("POST", "/api/widgets/([^/]+)", apiUpdateWidget(r.srv)),
	}
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type ctxKey struct{}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range rt.routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}