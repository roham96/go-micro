package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/roham96/go-micro/00_template/service"
	"github.com/roham96/go-micro/00_template/web"
)

func main() {
	fmt.Println("main()")
	router := web.NewRouter(service.New())
	log.Fatal(http.ListenAndServe("8080", router))
}