package gapi

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLogger(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	tm := time.Now()
	result, err := handler(ctx, req)
	code := codes.Unknown
	if stat, ok := status.FromError(err); ok {
		code = stat.Code()
	}
	if code != codes.OK {
		logrus.WithFields(logrus.Fields{
			"method":      info.FullMethod,
			"duration":    time.Since(tm),
			"status_code": code,
		}).Errorln("received a request")
	} else {
		logrus.WithFields(logrus.Fields{
			"method":      info.FullMethod,
			"duration":    time.Since(tm),
			"status_code": code,
		}).Infoln("received a request")
	}
	return result, err
}
