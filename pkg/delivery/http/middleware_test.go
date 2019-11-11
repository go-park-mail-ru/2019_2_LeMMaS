package http

import (
	testMock "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsOriginAllowed(t *testing.T) {
	m := newTestMiddleware(t)
	assert.False(t, m.isOriginAllowed(""))
	assert.True(t, m.isOriginAllowed("https://20192lemmasnew-h0ivhao1a.now.sh"))
	assert.False(t, m.isOriginAllowed("https://yandex.ru"))
	assert.False(t, m.isOriginAllowed("https://random324.now.sh"))
	assert.True(t, m.isOriginAllowed("http://localhost:3000"))
}

func newTestMiddleware(t *testing.T) Middleware {
	e := echo.New()
	logger := testMock.NewMockLogger(t)
	return NewMiddleware(e, logger)
}
