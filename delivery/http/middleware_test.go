package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsOriginAllowed(t *testing.T) {
	assert.False(t, IsOriginAllowed(""))
	assert.True(t, IsOriginAllowed("https://20192lemmasnew-h0ivhao1a.now.sh"))
	assert.False(t, IsOriginAllowed("https://yandex.ru"))
	assert.False(t, IsOriginAllowed("https://random324.now.sh"))
	assert.True(t, IsOriginAllowed("http://localhost:3000"))
}
