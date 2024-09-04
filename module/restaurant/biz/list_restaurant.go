package restaurantbiz

import (
	"context"
	"server/common"
	restaurantmodel "server/module/restaurant/model"

	"go.opencensus.io/trace"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

// type LikeRestaurantStore interface {
// 	GetRestautrantLikes(ctx context.Context, ids []int) (map[int]int, error)
// }

type listRestaurantBiz struct {
	store ListRestaurantStore
	// likeStore LikeRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	ctx1, span := trace.StartSpan(context, "biz.list_restaurant")
	result, err := biz.store.ListDataWithCondition(ctx1, filter, paging, "User")
	span.End()
	if err != nil {
		return nil, err
	}

	// ids := make([]int, len(result))

	// for i := range ids {
	// 	ids[i] = result[i].Id
	// }

	// likeMap, err := biz.likeStore.GetRestautrantLikes(context, ids)

	// if err != nil {
	// 	log.Println(err)
	// 	return result, nil
	// }

	// for i, item := range result {
	// 	result[i].LikedCount = likeMap[item.Id]
	// }

	return result, nil
}
