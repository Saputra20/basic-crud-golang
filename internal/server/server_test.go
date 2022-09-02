package server

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	s := NewServer()

	s.POST("/api", func(c *gin.Context) {})

	assertions := assert.New(t)

	assertions.NotZero(len(s.Routes()), "Routes should not zero")
}
