package ginrestaurant

import (
	"net/http"
	"server/common"
	"server/components/appctx"
	restaurantbiz "server/module/restaurant/biz"
	restaurantmodel "server/module/restaurant/model"
	restaurantstorage "server/module/restaurant/storage"

	// restaurantlikestorage "server/module/restaurantlike/storage"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filterData restaurantmodel.Filter
		if err := c.ShouldBind(&filterData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		// likestore := restaurantlikestorage.NewSQLStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filterData, &pagingData)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filterData))
	}
}
