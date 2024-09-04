package ginrestaurant

import (
	"net/http"
	"server/common"
	"server/components/appctx"
	restaurantbiz "server/module/restaurant/biz"
	restaurantstorage "server/module/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func FindRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewFindRestaurantBiz(store)

		data, err := biz.FindRestaurant(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
