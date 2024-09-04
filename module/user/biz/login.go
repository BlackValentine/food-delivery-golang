package userbiz

import (
	"context"
	"server/common"

	"server/components/tokenprovider"
	usermodel "server/module/user/model"
)

type LoginStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*usermodel.User, error)
}

type loginBusiness struct {
	storeUser          LoginStorage
	tokenProvider      tokenprovider.Provider
	hasher             Hasher
	expiryAccessToken  int
	expiryRefreshToken int
}

func NewLoginBusiness(
	storeUser LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	expiryAccessToken int,
	expiryRefreshToken int,
) *loginBusiness {
	return &loginBusiness{
		storeUser:          storeUser,
		tokenProvider:      tokenProvider,
		hasher:             hasher,
		expiryAccessToken:  expiryAccessToken,
		expiryRefreshToken: expiryRefreshToken,
	}
}

// 1. Find User
// 2. Has password from input and compare with password in db
// 3. Provider: issue JWT token for client
// 3.1. AccessToken and refreshToken
// 4. Return token(s)

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, *tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passHasher := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHasher {
		return nil, nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiryAccessToken)
	if err != nil {
		return nil, nil, common.ErrInternal(err)
	}

	refreshToken, err := business.tokenProvider.Generate(payload, business.expiryRefreshToken)
	if err != nil {
		return nil, nil, common.ErrInternal(err)
	}

	return accessToken, refreshToken, nil
}
