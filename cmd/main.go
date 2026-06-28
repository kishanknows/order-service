package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kishanknows/order-service/internal/config"
	"github.com/kishanknows/order-service/internal/database"
	"github.com/kishanknows/order-service/internal/routes"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	defer database.DB.Close()

	r := gin.Default()
	routes.RegisterOrderRoutes(r)

	server := &http.Server{
		Addr: config.Conf.Server.Addr,
		Handler: r,
		ReadTimeout: config.Conf.Server.ReadTimeout,
		WriteTimeout: config.Conf.Server.WriteTimeout,
		IdleTimeout: config.Conf.Server.IdleTimeout,
	}

	server.ListenAndServe()
}