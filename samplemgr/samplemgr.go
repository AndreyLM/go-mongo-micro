package main

import (
	"fmt"
	"strconv"

	"github.com/andreylm/go-mongo-micro/sqmplemgr/service"
	"github.com/urfave/negroni"
)

func main() {
	router := service.InitRouter()

	server := negroni.Classic()
	server.UseHandler(router)

	port := 33001
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	server.Run(addr)
}
