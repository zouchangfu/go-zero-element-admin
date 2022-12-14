package user

import (
	"github.com/zouchangfu/go-zero-element-admin/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/logic/user"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/types"
)

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserEditReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewEditLogic(r.Context(), svcCtx)
		err := l.Edit(&req)
		result.HttpResult(w, r, nil, err)
	}
}