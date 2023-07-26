package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/payment-system/portal-service/internal/config"
	"github.com/walbety/payment-system/portal-service/internal/rest"
	"github.com/walbety/payment-system/portal-service/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := config.Initialize()
	if err != nil {
		log.Fatal("error at initializing configs")
		os.Exit(2)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGINT)


	svc := service.New()

	log.Infof("Portal-service starting at port: %s", config.Env.FrontPort)
	log.Error("testando error")
	log.Debug("testando debug")
	log.Warn("teste de WARNING!!")
	log.WithFields(log.Fields{"field1" : "value1", "field2": "value aaaaa"}).Info("teste de fields")
	go func() {
		if err := rest.Start(svc); err != nil {
			log.WithError(err).Panic("error on http server")
		}
	}()

	<-stop
	rest.Stop(ctx)

	fmt.Print("aaaa\n\n")
}
