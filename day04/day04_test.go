package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordMeetsCriteria(t *testing.T) {
	assert.Equal(t, true, PasswordMeetsCriteria(11111, false))
	assert.Equal(t, false, PasswordMeetsCriteria(223450, false))
	assert.Equal(t, false, PasswordMeetsCriteria(123789, false))

	assert.Equal(t, true, PasswordMeetsCriteria(112233, true))
	assert.Equal(t, false, PasswordMeetsCriteria(123444, true))
	assert.Equal(t, true, PasswordMeetsCriteria(111122, true))
}

func TestGenerateAndCountPasswords(t *testing.T) {
	assert.Equal(t, 1955, GenerateAndCountPasswords(134792, 675810, false))
	assert.Equal(t, 1319, GenerateAndCountPasswords(134792, 675810, true))
}
