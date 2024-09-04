package ginuser

import (
	"net/http"
	"server/common"
	"server/components/appctx"
	"server/components/hasher"
	"server/components/tokenprovider/jwt"
	userbiz "server/module/user/biz"
	usermodel "server/module/user/model"
	userstorage "server/module/user/storage"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		tokenProvider := jwt.NewTokenJwtProvider(appCtx.SecretKey())
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60, 60*60*24)

		accessToken, _, err := biz.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(accessToken))
	}
}
