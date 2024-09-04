package restaurantbiz

import (
	"context"
	"server/common"
	restaurantmodel "server/module/restaurant/model"
)

type UpdateRestaurantStore interface {
	UpdateRestaurant(
		ctx context.Context,
		condition map[string]interface{},
		dataUpdate *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(context context.Context, id int, dataUpdate *restaurantmodel.RestaurantUpdate) error {
	if err := biz.store.UpdateRestaurant(context, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
