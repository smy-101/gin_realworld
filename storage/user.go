package storage

import (
	"context"
	"gin_realworld/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	_, err := db.ExecContext(ctx, "insert into user(username,password,email,image,bio) values(?,?,?,?,?)", user.UserName, user.Password, user.Email, user.Image, user.Bio)
	return err
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := db.GetContext(ctx, &user, "select username, password, email, image, bio from user where email = ?", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := db.GetContext(ctx, &user, "select username, password, email, image, bio from user where username = ?", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := db.ExecContext(ctx, "delete from user where email = ?", email)
	return err
}

func UpdateUserByUserName(ctx context.Context, username string, user *models.User) error {
	if user.Password != "" {
		_, err := db.ExecContext(ctx, "update user set username=?, password=?, email=?, image=?, bio=? where username=?",
			user.UserName, user.Password, user.Email, user.Image, user.Bio, username)
		return err
	}
	// if password is empty, do not update password
	_, err := db.ExecContext(ctx, "update user set username=?, email=?, image=?, bio=? where username=?",
		user.UserName, user.Email, user.Image, user.Bio, username)
	return err
}
