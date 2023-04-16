package handlers

import (
	"be/cmd/comment/pack"
	"be/cmd/comment/service"
	"be/grpc/commentdemo"
	"be/pkg/check"
	"be/pkg/errno"
	"context"
)

// // implements the service interface defined in IDL
type CommentServiceImpl struct {
	commentdemo.UnimplementedCommentServiceServer
}

// 创建评论或是 reply
// 如果 master 不为 0 则为 reply
func (s *CommentServiceImpl) CreateComment(ctx context.Context, req *commentdemo.CreateCommentRequest) (*commentdemo.CreateCommentResponse, error) {
	resp := new(commentdemo.CreateCommentResponse)

	// 检测参数
	if !check.CheckUserName(req.UserName) || !check.CheckCommentText(req.CommentText) || !check.CheckPostiveNumber(req.TargetID) || req.Master < 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	ids, err := service.NewCommentService(ctx).CreateComment(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.IDs = ids
	return resp, nil
}

// // 根据 ID 更新评论 不更新 评论者和被评论的文章
// func (s *CommentServiceImpl) UpdateComment(ctx context.Context, req *commentdemo.UpdateCommentRequest) (*commentdemo.UpdateCommentResponse, error) {
// 	resp := new(commentdemo.UpdateCommentResponse)

// 	// 检测参数
// 	if !check.CheckPostiveNumber(req.CommentID) {
// 		resp.Resp = pack.BuildResp(errno.ParamErr)
// 		return resp, nil
// 	}

// 	err := service.NewCommentService(context.Background()).UpdateComment(&commentdemo.UpdateCommentRequest{
// 		CommentID:   req.CommentID,
// 		CommentText: req.CommentText,
// 	})
// 	if err != nil {
// 		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
// 		return resp, nil
// 	}

// 	resp.Resp = pack.BuildResp(errno.Success)
// 	return resp, nil
// }

// 根据 ID 删除评论及其所有 reply
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *commentdemo.DeleteCommentRequest) (*commentdemo.DeleteCommentResponse, error) {
	resp := new(commentdemo.DeleteCommentResponse)

	// 检测参数
	if !check.CheckPostiveNumber(req.CommentID) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewCommentService(context.Background()).DeleteComment(&commentdemo.DeleteCommentRequest{
		CommentID: req.CommentID,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 根据 ID 查询评论 返回该评论及其所有回复
func (s *CommentServiceImpl) QueryComment(ctx context.Context, req *commentdemo.QueryCommentRequest) (*commentdemo.QueryCommentResponse, error) {
	resp := new(commentdemo.QueryCommentResponse)

	// CommentID 为空
	if len(req.CommentID) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	cms, err := service.NewCommentService(ctx).QueryComment(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	for _, cm := range cms {
		reply := make([]int32, 0)
		for _, rp := range cm.Reply {
			// // reply 绝对有 master
			// if rp.Master == nil {
			// 	resp.Resp = pack.BuildResp(errno.ConvertErr(err))
			// 	return resp, nil
			// }
			// reply = append(reply, &commentdemo.Comment{
			// 	ID:          int32(rp.ID),
			// 	TargetID:    int32(rp.TargetID),
			// 	UserName:    rp.UserName,
			// 	CommentText: rp.CommentText,
			// 	CreatedAt:   rp.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			// 	Master:      int32(*rp.Master),
			// 	Type:        rp.Type,
			// })
			reply = append(reply, int32(rp.ID))
		}

		var temp int32 = 0
		// 如果查询的目标是 reply
		if cm.Master != nil {
			temp = int32(*cm.Master)
		}

		resp.Comment = append(resp.Comment, &commentdemo.Comment{
			ID:          int32(cm.ID),
			TargetID:    int32(cm.TargetID),
			UserName:    cm.UserName,
			CommentText: cm.CommentText,
			CreatedAt:   cm.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),

			Master: temp,
			Reply:  reply,
			Type:   cm.Type,
		})
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

func (s *CommentServiceImpl) QueryCommentByTargetID(ctx context.Context, req *commentdemo.QueryCommentByTargetIDRequest) (*commentdemo.QueryCommentByTargetIDResponse, error) {
	resp := new(commentdemo.QueryCommentByTargetIDResponse)

	// 检测参数
	if !check.CheckPostiveNumber(req.TargetID) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	ids, err := service.NewCommentService(ctx).QueryCommentByTargetID(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.IDs = ids

	return resp, nil
}
