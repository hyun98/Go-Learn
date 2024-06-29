package server

import (
	"RPC-Learn/config"
	"RPC-Learn/gRPC/paseto"
	auth "RPC-Learn/gRPC/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type GRPCServer struct {
	auth.UnimplementedAuthServiceServer
	pasetoMaker    *paseto.PasetoMaker
	tokenVerifyMap map[string]*auth.AuthData
}

func NewGRPCServer(config *config.Config) error {
	if lis, err := net.Listen("tcp", config.GRPC.URL); err != nil {
		return err
	} else {
		server := grpc.NewServer([]grpc.ServerOption{}...)
		grpcServer := &GRPCServer{
			pasetoMaker:    paseto.NewPasetoMaker(config),
			tokenVerifyMap: make(map[string]*auth.AuthData),
		}

		auth.RegisterAuthServiceServer(server, grpcServer)
		reflection.Register(server)

		go func() {
			log.Println("Start GRPC Server")
			if err = server.Serve(lis); err != nil {
				panic(err)
			}
		}()
	}

	return nil
}

func (g *GRPCServer) CreateAuth(ctx context.Context, req *auth.CreateTokenReq) (*auth.CreateTokenRes, error) {
	data := req.Auth
	token := data.Token

	g.tokenVerifyMap[token] = data

	return &auth.CreateTokenRes{Auth: data}, nil
}

func (g *GRPCServer) VerifyAuth(ctx context.Context, req *auth.VerifyTokenReq) (*auth.VerifyTokenRes, error) {
	token := req.Token
	res := &auth.VerifyTokenRes{V: &auth.Verify{
		Auth: nil,
	}}

	if authData, ok := g.tokenVerifyMap[token]; !ok {
		res.V.Status = auth.ResponseType_FAILED
	} else if authData.ExpireDate < time.Now().Unix() {
		res.V.Status = auth.ResponseType_EXPIRED_DATE
	} else {
		res.V.Status = auth.ResponseType_SUCCESS
	}

	return res, nil
}
