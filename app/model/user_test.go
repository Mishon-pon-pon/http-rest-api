package model_test

import "testing"

import "github.com/mishon-pon-pon/http-rest-api/app/model"

import "github.com/stretchr/testify/assert"

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
