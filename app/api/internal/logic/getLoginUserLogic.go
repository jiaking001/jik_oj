package logic

import (
	"context"

	"app/api/internal/svc"
	"app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户登录信息
func NewGetLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginUserLogic {
	return &GetLoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginUserLogic) GetLoginUser() (resp *types.GetLoginUserResp, err error) {
	// todo: 通过rpc实现从数据库读取数据
	resp = &types.GetLoginUserResp{
		CreateTime:  "234",
		Id:          2,
		UpdateTime:  "234",
		UserAvatar:  "234",
		UserName:    "jiaking",
		UserProfile: "234",
		UserRole:    "admin",
	}
	return
}
