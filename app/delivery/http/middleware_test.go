package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsOriginAllowed(t *testing.T) {
	assert.False(t, isOriginAllowed(""))
	assert.True(t, isOriginAllowed("https://20192lemmasnew-h0ivhao1a.now.sh"))
	assert.False(t, isOriginAllowed("https://yandex.ru"))
	assert.False(t, isOriginAllowed("https://random324.now.sh"))
	assert.True(t, isOriginAllowed("http://localhost:3000"))
}
