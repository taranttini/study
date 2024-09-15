package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CepInvalidValue(t *testing.T) {

	result := NewValidateCep("2222-222")
	assert.False(t, result)

	result = NewValidateCep("2222a222")
	assert.False(t, result)
}

func Test_CepInvalidLength(t *testing.T) {

	result := NewValidateCep("22")
	assert.False(t, result)

	result = NewValidateCep("123456789")
	assert.False(t, result)
}

func Test_CepValidValue(t *testing.T) {
	result := NewValidateCep("12345123")
	assert.True(t, result)
}
