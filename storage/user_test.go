package storage

import (
	"context"
	"gin_realworld/models"
	"gin_realworld/utils"
	"testing"
)

func TestCreateAndGetUser(t *testing.T) {
	ctx := context.Background()
	userName := "xxx33x12"
	email := "xxx33123@gmail.com"

	err := CreateUser(ctx, &models.User{
		UserName: userName,
		Password: "xxxx123",
		Email:    email,
		Image:    "xxxx123",
		Bio:      "xxxx123",
	})
	if err != nil {
		t.Errorf("create user failed, err: %v", err)
		return
	}

	dbUser, err := GetUserByEmail(ctx, email)
	if err != nil {
		t.Errorf("get user by email failed, err: %v", err)
		return
	}

	t.Logf("get user: %v\n", utils.JsonMarshal(dbUser))

	err = DeleteUserByEmail(ctx, email)
	if err != nil {
		t.Errorf("delete user by email failed")
		return
	}

}
