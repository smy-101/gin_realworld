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

func DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := db.ExecContext(ctx, "delete from user where email = ?", email)
	return err
}
