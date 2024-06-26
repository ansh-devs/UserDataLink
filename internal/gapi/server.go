package gapi

import (
	"net"
	"os"

	"github.com/ansh-devs/userdatalink/internal/database"
	"github.com/ansh-devs/userdatalink/protos"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// UserService Definition of the UserService
type UserService struct {
	protos.UnimplementedUserServiceServer
	Repository database.UserRepository
	Srvr       *grpc.Server
	Port       string
	Ln         net.Listener
}

// NewUserService create an instance of User Service listens to the port and binds the listener to the UserService
func NewUserService() *UserService {
	srv := UserService{}

	srv.Port = ":" + os.Getenv("PORT")
	srv.Ln = MustListenToAddr(srv.Port)
	srv.Repository = database.NewUserRepository()

	srv.ConfigureGapi()

	return &srv
}

// ListenToAddr Attempts to listen to the port and return err if cannot
func MustListenToAddr(port string) net.Listener {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Fatalf("error while listening the port %v", err)
	}
	logrus.WithFields(logrus.Fields{"message": "listening on port " + port}).Info("user_service")
	return ln
}

func (s *UserService) ConfigureGapi() {
	grpcLogger := grpc.UnaryInterceptor(GrpcLogger)
	s.Srvr = grpc.NewServer(grpcLogger)
	protos.RegisterUserServiceServer(s.Srvr, s)
	reflection.Register(s.Srvr)
	logrus.WithFields(logrus.Fields{"message": "ready to accept connections"}).Info("user_service")

}

// mustEmbedUnimplementedUserServiceServer implements protos.UserServiceServer.
func (b *UserService) mustEmbedUnimplementedUserServiceServer() {}
