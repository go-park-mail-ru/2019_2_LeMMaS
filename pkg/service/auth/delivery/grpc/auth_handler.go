package grpc

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type AuthHandler struct {
	usecase auth.AuthUsecase
	server  *grpc.Server
}

func NewAuthHandler(usecase auth.AuthUsecase) *AuthHandler {
	h := AuthHandler{usecase: usecase}
	h.server = grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))
	auth.RegisterAuthServer(h.server, &h)
	return &h
}

func (h *AuthHandler) Serve(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	fmt.Println("listening " + address)
	return h.server.Serve(listener)
}

func (h *AuthHandler) Login(ctx context.Context, params *auth.LoginParams) (result *auth.LoginResult, grpcErr error) {
	result = &auth.LoginResult{}
	session, err := h.usecase.Login(params.Email, params.Password)
	if err != nil {
		result.Error = err.Error()
		return
	}
	result.Session = session
	return
}

func (h *AuthHandler) Logout(ctx context.Context, params *auth.LogoutParams) (result *auth.LogoutResult, grpcErr error) {
	result = &auth.LogoutResult{}
	err := h.usecase.Logout(params.Session)
	if err != nil {
		result.Error = err.Error()
	}
	return
}

func (h *AuthHandler) Register(ctx context.Context, params *auth.RegisterParams) (result *auth.RegisterResult, grpcErr error) {
	result = &auth.RegisterResult{}
	err := h.usecase.Register(params.Email, params.Password, params.Name)
	if err != nil {
		result.Error = err.Error()
	}
	return
}

func (h *AuthHandler) GetUser(ctx context.Context, params *auth.GetUserParams) (result *auth.GetUserResult, grpcErr error) {
	result = &auth.GetUserResult{}
	id, ok := h.usecase.GetUser(params.Session)
	if !ok {
		result.Error = consts.ErrNotFound.Error()
		return
	}
	result.Id = int32(id)
	return
}

func (h *AuthHandler) GetPasswordHash(ctx context.Context, params *auth.GetPasswordHashParams) (*auth.GetPasswordHashResult, error) {
	hash := h.usecase.GetPasswordHash(params.Password)
	return &auth.GetPasswordHashResult{PasswordHash: hash}, nil
}
