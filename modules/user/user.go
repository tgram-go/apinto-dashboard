package user

import (
	"context"
	user_dto "github.com/eolinker/apinto-dashboard/modules/user/user-dto"
	"github.com/eolinker/apinto-dashboard/modules/user/user-model"
)

type IUserInfoService interface {
	GetAllUsers(ctx context.Context) ([]*user_model.UserInfo, error)
	GetUserInfo(ctx context.Context, userId int) (*user_model.UserInfo, error)
	GetUserInfoMaps(ctx context.Context, userId ...int) (map[int]*user_model.UserInfo, error)
	GetUserInfoByName(ctx context.Context, userName string) (*user_model.UserInfo, error)
	GetUserInfoByNames(ctx context.Context, userNames ...string) (map[string]*user_model.UserInfo, error)
	UpdateMyProfile(ctx context.Context, userId int, req *user_dto.UpdateMyProfileReq) error
}
