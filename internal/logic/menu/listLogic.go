package menu

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/errx"

	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() (resp *types.MenuListResp, err error) {
	allMenus, sqlResult := l.svcCtx.MenuDao.ListAll()
	if sqlResult.Error != nil {
		return nil, errx.NewErrCode(errx.DbError)
	}

	var menus []*types.MenuResp
	for _, u := range allMenus {
		menuResp := types.MenuResp{}
		copier.Copy(&menuResp, &u)
		menuResp.CreatedAt = u.CreatedAt.UnixMilli()
		menuResp.UpdatedAt = u.UpdatedAt.UnixMilli()
		menus = append(menus, &menuResp)
	}

	return &types.MenuListResp{MenuList: menus}, nil
}
