package app

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCheckAuth(t *testing.T) {
	s := "Basic dXNlcjE6cHdkMQ=="

	userPwd := basicAuth(s)
	assert.Equal(t, "user1", userPwd[0])
	assert.Equal(t, "pwd1", userPwd[1])
}
