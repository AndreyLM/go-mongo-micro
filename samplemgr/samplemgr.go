package main

import (
	"fmt"

	"github.com/andreylm/go-mongo-micro/sqmplemgr/config"
	"github.com/andreylm/go-mongo-micro/sqmplemgr/db"
	"github.com/andreylm/go-mongo-micro/sqmplemgr/logger"

	"github.com/andreylm/go-mongo-micro/sqmplemgr/service"
	"github.com/urfave/negroni"
)

func main() {
	config.Load()

	logger.Init()
	db.Init()

	router := service.InitRouter()

	server := negroni.Classic()
	server.UseHandler(router)

	port := config.ServicePort()
	addr := fmt.Sprintf(":%s", port)

	server.Run(addr)
}
