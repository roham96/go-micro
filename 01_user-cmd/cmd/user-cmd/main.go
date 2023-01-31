package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/roham96/go-micro/01_user-cmd/service"
	"github.com/roham96/go-micro/01_user-cmd/web"
)

func main() {
	fmt.Println("main()")
	router := web.NewRouter(service.New())
	log.Fatal(http.ListenAndServe("8080", router))
}