package gapi

import (
	"context"
	"net"
	"testing"

	"github.com/ansh-devs/userdatalink/internal/database"
	"github.com/ansh-devs/userdatalink/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

func TestUserService_GetUsersListByIds(t *testing.T) {
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
	resp, err := client.GetUsersListByIds(context.Background(), &protos.GetUsersListByIdsRequest{Ids: []int64{5, 8}})
	if err != nil {
		t.Fatalf("GetUsersListByIds error %v", err)
	}
	for _, v := range resp.GetUsers() {
		if !(v.Id == 5 || v.Id == 8) {
			t.Fatalf("Unexpected values %v", v)
		}
	}

	// Case when out of bound id is passed resulting in InvalidArgument proto error
	_, err = client.GetUsersListByIds(context.Background(), &protos.GetUsersListByIdsRequest{Ids: []int64{42}})
	code, _ := status.FromError(err)
	if code.Code() != codes.InvalidArgument {
		t.Fatalf("GetUsersListByIds error Invalid Argument unexpected code %v", err)
	}
}
