package http

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMetricsHandler(t *testing.T) {
	assert.NotNil(t, NewMetricsHandler(echo.New()))
}
