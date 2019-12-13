package user

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user/proto"
	"golang.org/x/net/context"
)

type Usecase interface {
	GetUserBySessionID(context.Context, *user.SessionID) (*user.UserAndError, error)
	UpdateUser(context.Context, *user.UserToUpdate) (*user.Error, error)
	UpdateUserAvatar(context.Context, *user.UserToUpdateAvatar) (*user.Error, error)
	GetLeaderUsers(context.Context, *user.UserID) (*user.Users, error)
	GetUserByID(context.Context, *user.UserID) (*user.UserAndError, error)
	GetAvatarUrlByName(context.Context, *user.UserName) (*user.AvatarUrl, error)
}
