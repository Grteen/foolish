package handlers

type UserParma struct {
	Name     string `form:"name"`
	PassWord string `form:"password"`
	Email    string `form:"email"`
}

type LoginParma struct {
	NameOrEmail string `form:"account"`
	PassWord    string `form:"password"`
}

type UpdateUserInfoParma struct {
	UserName    string `form:"username"`
	NickName    string `form:"nickname"`
	Description string `form:"description"`
	Avator      string `form:"avator"`
}

type QueryUserInfoParma struct {
	UserName string `form:"username"`
}

type UploadPicParma struct {
	UserName string `form:"username"`
	PicName  string `form:"picname"`
}
