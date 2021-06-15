package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockUser struct{}

func (user *mockUser) FirstLastName() string {
	return "Mocking Jay"
}

func TestUser_FirstLastName(t *testing.T) {
	client := &mockUser{}

	length := GetFirstLastNameLength(client)
	assert.Equal(t, 11, length, "Two length should be the same")
}
