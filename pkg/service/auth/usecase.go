package auth

type AuthUsecase interface {
	Login(email, password string) (sessionID string, err error)
	Logout(sessionID string) error
	Register(email, password, name string) error
}
