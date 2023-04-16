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

func QueryArticalEx(ctx context.Context, req *articaldemo.QueryArticalRequest) ([]*articaldemo.Artical, error) {
	resp, err := articalClient.QueryArticalEx(ctx, req)
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

func QueryLikeStar(ctx context.Context, req *articaldemo.QueryLikeStarRequest) (*articaldemo.LikeStar, error) {
	resp, err := articalClient.QueryLikeStar(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.LikeStar, nil
}

func QueryAllLikeStar(ctx context.Context, req *articaldemo.QueryAllLikeStarRequest) ([]int32, error) {
	resp, err := articalClient.QueryAllLikeStar(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.ArticalIDs, nil
}

func RdbSetArtical(ctx context.Context, req *articaldemo.RdbSetArticalRequest) error {
	resp, err := articalClient.RdbSetArtical(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func RdbDelArtical(ctx context.Context, req *articaldemo.RdbDelArticalRequest) error {
	resp, err := articalClient.RdbDelArtical(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func RdbGetArtical(ctx context.Context, req *articaldemo.RdbGetArticalRequest) ([]*articaldemo.RdbArtical, []int32, error) {
	resp, err := articalClient.RdbGetArtical(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.RdbArticals, resp.Ungot, nil
}

// 不获取 Text 的版本
func RdbGetArticalEx(ctx context.Context, req *articaldemo.RdbGetArticalRequest) ([]*articaldemo.RdbArtical, []int32, error) {
	resp, err := articalClient.RdbGetArticalEx(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.RdbArticals, resp.Ungot, nil
}

func RdbIncreaseitf(ctx context.Context, req *articaldemo.RdbIncreaseitfRequest) error {
	resp, err := articalClient.RdbIncreaseitf(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func CreateStar(ctx context.Context, req *articaldemo.CreateStarRequest) error {
	resp, err := articalClient.CreateStar(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func CreateStarFolder(ctx context.Context, req *articaldemo.CreateStarFolderRequest) error {
	resp, err := articalClient.CreateStarFolder(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func DeleteStarFolderAndMove(ctx context.Context, req *articaldemo.DeleteStarFolderAndMoveRequest) error {
	resp, err := articalClient.DeleteStarFolderAndMove(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func UpdateStarFolder(ctx context.Context, req *articaldemo.UpdateStarFolderRequest) error {
	resp, err := articalClient.UpdateStarFolder(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func QueryStarFolder(ctx context.Context, req *articaldemo.QueryStarFolderRequest) ([]*articaldemo.StarFolder, error) {
	resp, err := articalClient.QueryStarFolder(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.StarFolders, nil
}

func QueryAllStarFolder(ctx context.Context, req *articaldemo.QueryAllStarFolderRequest) ([]*articaldemo.StarFolder, error) {
	resp, err := articalClient.QueryAllStarFolder(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.StarFolders, nil
}

func QueryAllStar(ctx context.Context, req *articaldemo.QueryAllStarRequest) ([]*articaldemo.Star, error) {
	resp, err := articalClient.QueryAllStar(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Stars, nil
}

func UpdateStarOwner(ctx context.Context, req *articaldemo.UpdateStarOwnerRequest) error {
	resp, err := articalClient.UpdateStarOwner(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}
