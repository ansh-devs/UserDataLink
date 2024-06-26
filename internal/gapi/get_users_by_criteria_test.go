package gapi

import (
	"context"
	"net"
	"testing"

	"github.com/ansh-devs/userdatalink/internal/database"
	"github.com/ansh-devs/userdatalink/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestUserService_GetUsersByCriteria(t *testing.T) {
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
	resp, err := client.GetUsersByCriteria(context.Background(), &protos.GetUsersByCriteriaRequest{Type: 2, Value: "kalkara"})
	if err != nil {
		t.Fatalf("GetUsersByCriteria error %v", err)
	}
	for _, v := range resp.GetUsers() {
		if !(v.Id == 8 && v.FName == "Clevey Fazzioli") {
			t.Fatalf("Unexpected values %v", v)
		}
	}
}
