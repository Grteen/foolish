package handlers

import (
	"be/cmd/api/pack"
	"be/pkg/constants"
	"be/pkg/errno"

	"github.com/gin-gonic/gin"
)

// 上传图片，根据 userName 来确认存储的文件夹
func UploadPic(ctx *gin.Context) {
	var p UploadPicParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 账户 图片名字 为空
	if len(p.UserName) == 0 || len(p.PicName) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与账户相匹配
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 获取图片文件
	f, err := ctx.FormFile("pic")
	if err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 待做 优化字符串拼接
	picstr := constants.PicUploadDir + "/" + p.UserName + "/" + p.PicName + ".webp"
	err = ctx.SaveUploadedFile(f, picstr)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, gin.H{
		"uri": constants.PicHttpUri + "/" + p.UserName + "/" + p.PicName + ".webp",
	})
}
