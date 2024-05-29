package main

import (
	"github.com/ansh-devs/tc-assessment/internal/gapi"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
}

func main() {

	userService := gapi.NewUserService()
	defer userService.Srvr.GracefulStop()
	if err := userService.Srvr.Serve(userService.Ln); err != nil {
		logrus.Fatalf("cannot start gRPC server : %v", err)
	} else {

	}
}
