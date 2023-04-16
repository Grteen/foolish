package rpc

import (
	"be/grpc/searchdemo"
	"be/pkg/errno"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var searchClient searchdemo.SearchServiceClient

func InitSearchRpc() {
	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	searchClient = searchdemo.NewSearchServiceClient(conn)
}

func SearchArtical(ctx context.Context, req *searchdemo.SearchArticalRequest) ([]int32, error) {
	resp, err := searchClient.SearchArtical(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.ArticalIDs, err
}

func SearchUserZoom(ctx context.Context, req *searchdemo.SearchUserZoomRequest) ([]*searchdemo.Target, error) {
	resp, err := searchClient.SearchUserZoom(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Targets, nil
}
