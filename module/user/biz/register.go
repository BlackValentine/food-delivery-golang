package userbiz

import (
	"context"
	"errors"
	"server/common"
	usermodel "server/module/user/model"
)

type RegisterStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStore RegisterStorage
	hasher        Hasher
}

func NewRegisterBusiness(registerStore RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{registerStore: registerStore, hasher: hasher}
}

func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := business.registerStore.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		if user.Status == 0 {
			return errors.New("user has been disabled")
		}
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" //hard code

	if err := business.registerStore.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
