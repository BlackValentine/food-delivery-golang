package ginrestaurantlike

import (
	"net/http"
	"server/common"
	"server/components/appctx"
	restaurantstorage "server/module/restaurant/storage"
	restaurantlikebiz "server/module/restaurantlike/biz"
	restaurantlikemodel "server/module/restaurantlike/model"
	restaurantlikestorage "server/module/restaurantlike/storage"

	"github.com/gin-gonic/gin"
)

// POST /v1/restaurant/:id/like

func LikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(db)
		incStore := restaurantstorage.NewSQLStore(db)
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, incStore)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
