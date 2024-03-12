package data

import (
	"context"
	"hello-go/internal/biz"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (r *userRepo) Create(ctx context.Context, u *biz.User) (*biz.User, error) {
	if err := r.data.Db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &biz.User{ID: u.ID, Username: u.Username}, nil
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	var u biz.User
	ret := r.data.Db.Where("username = ?", username).First(&u)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &u, nil
}
