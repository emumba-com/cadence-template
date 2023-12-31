package main

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"
	"net/http"
	"orderService/clients/cadenceClient"
	"orderService/controllers/v1/orders"
	"orderService/env"
	"orderService/log"
	"time"
)

func main() {
	fileWriter := log.InitLogger(env.Env.BuildEnv, "./logs/order-service.log")

	logger := log.GetLogger()
	logger.Info("Starting Order Service")

	cadenceCli, err := cadenceClient.GetNewCadenceClient()
	if err != nil {
		logger.Error(err.Error())

		return
	}

	cadStore := cadenceClient.NewStore(&cadenceCli)

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: fileWriter,
	}))

	router.Use(apmgin.Middleware(router))

	httpClient := http.DefaultClient
	httpClient.Timeout = time.Minute * 5
	orderServiceGrp := router.Group("order-service/api/v1")

	orders.CreateNewServer(router, orderServiceGrp, cadStore)

	httpServer := &http.Server{
		Addr:    ":" + env.Env.ServerPort,
		Handler: router,
	}

	startHttpServer(httpServer)
}

func startHttpServer(server *http.Server) {
	logger := log.GetLogger()
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(err.Error())
	}
}
