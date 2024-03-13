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

// Create 创建用户
func (r *userRepo) Create(ctx context.Context, u *biz.User) (*biz.User, error) {
	if err := r.data.Db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &biz.User{ID: u.ID, Username: u.Username}, nil
}

// FindByUsername 根据username获取用户
func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	var u biz.User
	ret := r.data.Db.Where("username = ?", username).First(&u)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &u, nil
}

// FindById 根据ID获取用户
func (r *userRepo) FindById(ctx context.Context, id uint) (*biz.User, error) {
	var u biz.User
	ret := r.data.Db.Find(&u, id)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &u, nil
}
