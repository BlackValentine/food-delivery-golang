package ginrestaurantlike

import (
	"net/http"
	"server/common"
	"server/components/appctx"
	restaurantstorage "server/module/restaurant/storage"
	restaurantlikebiz "server/module/restaurantlike/biz"
	restaurantlikestorage "server/module/restaurantlike/storage"

	"github.com/gin-gonic/gin"
)

// POST /v1/restaurant/:id/unlike

func UnlikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		restaurantId := int(uid.GetLocalID())
		userId := requester.GetUserId()

		store := restaurantlikestorage.NewSQLStore(db)
		decStore := restaurantstorage.NewSQLStore(db)
		biz := restaurantlikebiz.NewUserUnlikeRestaurantBiz(store, decStore)

		if err := biz.UnlikeRestaurant(c.Request.Context(), userId, restaurantId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
