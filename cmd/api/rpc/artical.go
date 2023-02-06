package rpc

import (
	"be/grpc/articaldemo"
	"be/pkg/errno"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var articalClient articaldemo.ArticalServiceClient

func InitArticalRpc() {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	articalClient = articaldemo.NewArticalServiceClient(conn)
}

func CreateArtical(ctx context.Context, req *articaldemo.CreateArticalRequest) error {
	resp, err := articalClient.CreateArtical(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func QueryArtical(ctx context.Context, req *articaldemo.QueryArticalRequest) (*articaldemo.Artical, error) {
	resp, err := articalClient.QueryArtical(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Artical, err
}

func CreateLikeStar(ctx context.Context, req *articaldemo.CreateLikeStarRequest) error {
	resp, err := articalClient.CreateLikeStar(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func DeleteLikeStar(ctx context.Context, req *articaldemo.DeleteLikeStarRequest) error {
	resp, err := articalClient.DeleteLikeStar(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}
