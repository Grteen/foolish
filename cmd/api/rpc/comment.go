package rpc

import (
	"be/grpc/commentdemo"
	"be/pkg/errno"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var commentClient commentdemo.CommentServiceClient

func InitCommentRpc() {
	conn, err := grpc.Dial("127.0.0.1:8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	commentClient = commentdemo.NewCommentServiceClient(conn)
}

func CreateComment(ctx context.Context, req *commentdemo.CreateCommentRequest) ([]int32, error) {
	resp, err := commentClient.CreateComment(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.IDs, nil
}

func QueryComment(ctx context.Context, req *commentdemo.QueryCommentRequest) ([]*commentdemo.Comment, error) {
	resp, err := commentClient.QueryComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Comment, nil
}

func QueryCommentByTargetID(ctx context.Context, req *commentdemo.QueryCommentByTargetIDRequest) ([]int32, error) {
	resp, err := commentClient.QueryCommentByTargetID(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.IDs, nil
}

func DeleteComment(ctx context.Context, req *commentdemo.DeleteCommentRequest) error {
	resp, err := commentClient.DeleteComment(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func UpdateComment(ctx context.Context, req *commentdemo.UpdateCommentRequest) error {
	resp, err := commentClient.UpdateComment(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}
