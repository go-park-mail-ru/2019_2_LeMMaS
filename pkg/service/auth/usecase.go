//go:generate mockgen -source=$GOFILE -destination=usecase_mock.go -package=$GOPACKAGE

package auth

type AuthUsecase interface {
	Login(email, password string) (session string, err error)
	Logout(session string) error
	Register(email, password, name string) error
	GetUser(session string) (int, bool)
	GetPasswordHash(password string) string
}
