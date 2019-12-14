package grpc

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type UserHandler struct {
	usecase user.UserUsecase
	server  *grpc.Server
}

func NewUserHandler(usecase user.UserUsecase) *UserHandler {
	h := UserHandler{usecase: usecase}
	h.server = grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))
	user.RegisterUserServer(h.server, &h)
	return &h
}

func (h *UserHandler) Serve(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	fmt.Println("listening " + address)
	return h.server.Serve(listener)
}

func (h *UserHandler) GetAll(ctx context.Context, params *user.GetAllParams) (result *user.GetAllResult, grpcErr error) {
	result = &user.GetAllResult{}
	users, err := h.usecase.GetAll()
	if err != nil {
		result.Error = err.Error()
		return
	}
	result.Users = h.convertUsers(users)
	return
}

func (h *UserHandler) GetByID(ctx context.Context, params *user.GetByIDParams) (result *user.GetByIDResult, grpcErr error) {
	result = &user.GetByIDResult{}
	u, err := h.usecase.GetByID(int(params.Id))
	if err != nil {
		result.Error = err.Error()
		return
	}
	result.User = h.convertUser(u)
	return
}

func (h *UserHandler) GetByEmail(ctx context.Context, params *user.GetByEmailParams) (result *user.GetByEmailResult, grpcErr error) {
	result = &user.GetByEmailResult{}
	u, err := h.usecase.GetByEmail(params.Email)
	if err != nil {
		result.Error = err.Error()
		return
	}
	result.User = h.convertUser(u)
	return
}

func (h *UserHandler) Update(ctx context.Context, params *user.UpdateParams) (result *user.UpdateResult, grpcErr error) {
	result = &user.UpdateResult{}
	err := h.usecase.Update(int(params.Id), params.PasswordHash, params.Name)
	if err != nil {
		result.Error = err.Error()
	}
	return
}

func (h *UserHandler) UpdateAvatar(context.Context, *user.UpdateAvatarParams) (*user.UpdateAvatarResult, error) {
	return nil, nil
}

func (h *UserHandler) GetSpecialAvatar(context.Context, *user.GetSpecialAvatarParams) (*user.GetSpecialAvatarResult, error) {
	return nil, nil
}

func (h *UserHandler) convertUsers(users []*model.User) []*user.UserData {
	result := make([]*user.UserData, 0)
	for _, usr := range users {
		result = append(result, h.convertUser(usr))
	}
	return result
}

func (h *UserHandler) convertUser(usr *model.User) *user.UserData {
	return &user.UserData{
		Id:           int32(usr.ID),
		Email:        usr.Email,
		PasswordHash: usr.PasswordHash,
		Name:         usr.Name,
		AvatarPath:   usr.AvatarPath,
	}
}
