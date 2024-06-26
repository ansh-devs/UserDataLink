package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ansh-devs/userdatalink/internal/gapi"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
}

func main() {
	errs := make(chan error)
	userService := gapi.NewUserService()
	defer userService.Srvr.GracefulStop()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		errs <- userService.Srvr.Serve(userService.Ln)
	}()

	for sig := range errs {
		logrus.WithFields(logrus.Fields{
			"status":  sig,
			"message": "gracefully shutting down",
		}).Warn("user_service")
		os.Exit(0)
	}
}
