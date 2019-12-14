package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCsrfUsecase_TokenBySession(t *testing.T) {
	usecase := NewCSRFUsecase()

	session := "8386edb6-2523-444d-8c99-a959db3a60ab"
	token, err := usecase.CreateTokenBySession(session)
	assert.NotEmpty(t, token)
	assert.NoError(t, err)

	ok, err := usecase.CheckTokenBySession(token, session)
	assert.True(t, ok)
	assert.NoError(t, err)
}
