package access

type CsrfUsecase interface {
	CreateSimpleToken() (string, error)
	CheckSimpleToken(token string) (bool, error)
	CreateTokenBySession(sessionID string) (string, error)
	CheckTokenBySession(token string, sessionID string) (bool, error)
}
