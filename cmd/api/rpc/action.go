package rpc

import (
	"be/grpc/actiondemo"
	"be/pkg/errno"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var actionClient actiondemo.ActionServiceClient

func InitActionRpc() {
	conn, err := grpc.Dial("127.0.0.1:8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	actionClient = actiondemo.NewActionServiceClient(conn)
}

// 创建动态
func CreateAction(ctx context.Context, req *actiondemo.CreateActionRequest) error {
	resp, err := actionClient.CreateAction(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

// 删除动态
func DeleteAction(ctx context.Context, req *actiondemo.DeleteActionRequest) error {
	resp, err := actionClient.DeleteAction(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

// 根据ID查询动态
func QueryAction(ctx context.Context, req *actiondemo.QueryActionRequest) ([]*actiondemo.Action, error) {
	resp, err := actionClient.QueryAction(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Actions, nil
}

// 查询某人的所有动态
func QueryActionByAuthor(ctx context.Context, req *actiondemo.QueryActionByAuthorRequest) ([]int32, error) {
	resp, err := actionClient.QueryActionByAuthor(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.IDs, nil
}

// 点赞动态
func CreateActionLike(ctx context.Context, req *actiondemo.CreateActionLikeRequest) error {
	resp, err := actionClient.CreateActionLike(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

// 删除点赞
func DeleteActionLike(ctx context.Context, req *actiondemo.DeleteActionLikeRequest) error {
	resp, err := actionClient.DeleteActionLike(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

// 查询动态点赞
func QueryActionLike(ctx context.Context, req *actiondemo.QueryActionLikeRequest) ([]*actiondemo.ActionLike, error) {
	resp, err := actionClient.QueryActionLike(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Actionlikes, nil
}

func CreateActionComment(ctx context.Context, req *actiondemo.CreateCommentRequest) ([]int32, error) {
	resp, err := actionClient.CreateComment(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.IDs, nil
}

func QueryActionComment(ctx context.Context, req *actiondemo.QueryCommentRequest) ([]*actiondemo.Comment, error) {
	resp, err := actionClient.QueryComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Comment, nil
}

func QueryActionCommentByActionID(ctx context.Context, req *actiondemo.QueryCommentByActionIDRequest) ([]int32, error) {
	resp, err := actionClient.QueryCommentByActionID(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.IDs, nil
}

func DeleteActionComment(ctx context.Context, req *actiondemo.DeleteCommentRequest) error {
	resp, err := actionClient.DeleteComment(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func UpdateActionComment(ctx context.Context, req *actiondemo.UpdateCommentRequest) error {
	resp, err := actionClient.UpdateComment(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}
