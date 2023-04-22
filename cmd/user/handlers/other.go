package handlers

import (
	"be/cmd/user/dal/db"
	"be/cmd/user/dal/rdb"
	"be/cmd/user/pack"
	"be/grpc/userdemo"
)

func ChangeUserInfoToRdbUserInfo(ufs []*db.UserInfo) []*rdb.RdbUser {
	res := make([]*rdb.RdbUser, 0)
	for _, uf := range ufs {
		res = append(res, &rdb.RdbUser{
			UserName:    uf.UserName,
			NickName:    uf.NickName,
			Description: uf.Description,
			UserAvator:  uf.UserAvator,
		})
	}
	return res
}

func ChangeUserToRdbUser(us []*db.User) []*rdb.RdbUser {
	res := make([]*rdb.RdbUser, 0)
	for _, u := range us {
		res = append(res, &rdb.RdbUser{
			UserName:        u.UserName,
			SubNum:          u.SubNum,
			FanNum:          u.FanNum,
			ArtNum:          u.ArtNum,
			IsAdministrator: u.IsAdministrator,
			FanPublic:       u.FanPublic,
			SubPublic:       u.SubPublic,
		})
	}
	return res
}

func ChangeDbPubNoticeToRPCPubNotice(pubs []*db.PubNotice) []*userdemo.PubNotice {
	res := make([]*userdemo.PubNotice, 0)
	for _, p := range pubs {
		res = append(res, &userdemo.PubNotice{
			ID:        int32(p.ID),
			CreatedAt: p.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			UserName:  p.UserName,
			Text:      p.Text,
		})
	}
	return res
}
