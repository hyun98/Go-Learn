package client

import (
	"RPC-Learn/config"
	"RPC-Learn/gRPC/paseto"
	auth "RPC-Learn/gRPC/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type GRPCClient struct {
	client            *grpc.ClientConn
	authServiceClient auth.AuthServiceClient
	pasetoMaker       *paseto.PasetoMaker
}

func NewGRPCClient(config *config.Config) (*GRPCClient, error) {
	c := &GRPCClient{}

	if client, err := grpc.NewClient(config.GRPC.URL, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	} else {
		c.client = client
		c.authServiceClient = auth.NewAuthServiceClient(c.client)
		c.pasetoMaker = paseto.NewPasetoMaker(config)
	}

	return c, nil
}

func (g *GRPCClient) CreateAuth(name string) (*auth.CreateTokenRes, error) {
	now := time.Now()
	expiredTime := now.Add(30 * time.Minute)

	authData := &auth.AuthData{
		Name:       name,
		CreateDate: now.Unix(),
		ExpireDate: expiredTime.Unix(),
	}

	if token, err := g.pasetoMaker.CreateNewToken(authData); err != nil {
		return nil, err
	} else {
		authData.Token = token

		if res, err := g.authServiceClient.CreateAuth(context.Background(), &auth.CreateTokenReq{Auth: authData}); err != nil {
			return nil, err
		} else {
			return res, nil
		}
	}
}

func (g *GRPCClient) VerifyAuth(token string) (*auth.VerifyTokenRes, error) {
	if res, err := g.authServiceClient.VerifyAuth(context.Background(), &auth.VerifyTokenReq{Token: token}); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
