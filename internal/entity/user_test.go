package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jhon Doe", "J@j.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Jhon Doe", user.Name)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "J@j.com", user.Email)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("Jhon Doe", "J@j.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
