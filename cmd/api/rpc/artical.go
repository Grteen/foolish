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

func DeleteArtical(ctx context.Context, req *articaldemo.DeleteArticalRequest) error {
	resp, err := articalClient.DeleteArtical(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func UpdateArtical(ctx context.Context, req *articaldemo.UpdateArticalRequest) error {
	resp, err := articalClient.UpdateArtical(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func QueryArtical(ctx context.Context, req *articaldemo.QueryArticalRequest) ([]*articaldemo.Artical, error) {
	resp, err := articalClient.QueryArtical(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Artical, err
}

func QueryArticalByAuthor(ctx context.Context, req *articaldemo.QueryArticalByAuthorRequest) ([]int32, error) {
	resp, err := articalClient.QueryArticalByAuthor(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.IDs, err
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

func QueryAllLikeStar(ctx context.Context, req *articaldemo.QueryAllLikeStarRequest) ([]uint32, error) {
	resp, err := articalClient.QueryAllLikeStar(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.ArticalIDs, nil
}

func CreateComment(ctx context.Context, req *articaldemo.CreateCommentRequest) error {
	resp, err := articalClient.CreateComment(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func QueryComment(ctx context.Context, req *articaldemo.QueryCommentRequest) ([]*articaldemo.Comment, error) {
	resp, err := articalClient.QueryComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Comment, nil
}

func QueryCommentByArticalID(ctx context.Context, req *articaldemo.QueryCommentByArticalIDRequest) ([]int32, error) {
	resp, err := articalClient.QueryCommentByArticalID(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.IDs, nil
}

func DeleteComment(ctx context.Context, req *articaldemo.DeleteCommentRequest) error {
	resp, err := articalClient.DeleteComment(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func UpdateComment(ctx context.Context, req *articaldemo.UpdateCommentRequest) error {
	resp, err := articalClient.UpdateComment(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}
