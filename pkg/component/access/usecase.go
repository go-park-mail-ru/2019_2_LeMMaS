//go:generate mockgen -source=$GOFILE -destination=usecase_mock.go -package=$GOPACKAGE

package access

type CsrfUsecase interface {
	CreateTokenBySession(sessionID string) (string, error)
	CheckTokenBySession(token string, sessionID string) (bool, error)
}
