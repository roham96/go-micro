package web

import (
	"fmt"
	"net/http"

	"github.com/roham96/go-micro/01_user-cmd/service"
)

func apiGetWidgets(srv *service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		srv.ApiGetWidgets()
		fmt.Fprint(w, "ApiGetWidgets\n")
	}
}

func apiCreateWidget(srv *service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		srv.ApiCreateWidget()
		fmt.Fprint(w, "ApiCreateWidget\n")
	}
}

func apiUpdateWidget(srv *service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		srv.ApiUpdateWidget()
		fmt.Fprint(w, "ApiUpdateWidget\n")
	}
}