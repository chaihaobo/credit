package main

import (
	"credit-platform/infrastructure"
	"credit-platform/resource"
	"credit-platform/service"
	"credit-platform/transport"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	res := resource.New("config.yaml")
	infra := infrastructure.New(res)
	svc := service.New(infra, res)
	tp := transport.New(res, svc, infra)
	tp.Serve()

	// GRACEFUL SHUTDOWN
	graceful := make(chan os.Signal, 1)
	signal.Notify(graceful, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-graceful
	tp.Shutdown()

}
