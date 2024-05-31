package gapi

import (
	"context"
	"net"
	"testing"

	"github.com/ansh-devs/tc-assessment/internal/database"
	"github.com/ansh-devs/tc-assessment/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestUserService_GetuserById(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)

	// equivalent to defer
	t.Cleanup(func() {
		lis.Close()
	})
	srvr := grpc.NewServer()

	t.Cleanup(func() {
		srvr.Stop()
	})

	svc := UserService{
		Repository: database.NewUserRepository(),
	}
	protos.RegisterUserServiceServer(srvr, &svc)
	go func() {
		if err := srvr.Serve(lis); err != nil {
			t.Fatalf("srvr.Serve :%v", err)
		}
	}()

	// Testing
	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}
	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	t.Cleanup(func() {
		conn.Close()
	})
	if err != nil {
		t.Fatalf("grpc.DialContext %v", err)
	}
	client := protos.NewUserServiceClient(conn)
	resp, err := client.GetUserById(context.Background(), &protos.GetUserByIdRequest{Id: 3})
	if err != nil {
		t.Fatalf("GetUserById error %v", err)
	}
	if resp.GetUser().FName != "Nerita Bomfield" {
		t.Fatalf("Unexpected values %v", resp.User)
	}
}
