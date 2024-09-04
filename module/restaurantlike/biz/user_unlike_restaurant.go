package restaurantlikebiz

import (
	"context"
	"log"
	"server/common"
	restaurantlikemodel "server/module/restaurantlike/model"
)

type UserUnlikeRestaurantStore interface {
	Delete(
		ctx context.Context,
		userId, restaurantId int,
	) error
}

type DecLikedCountRestaurantStore interface {
	DecreaseLikedCount(ctx context.Context, id int) error
}

type userUnlikeRestaurantBiz struct {
	store    UserUnlikeRestaurantStore
	decStore DecLikedCountRestaurantStore
}

func NewUserUnlikeRestaurantBiz(
	store UserUnlikeRestaurantStore,
	decStore DecLikedCountRestaurantStore,
) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{
		store:    store,
		decStore: decStore,
	}
}

func (biz *userUnlikeRestaurantBiz) UnlikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotUnlikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		if err := biz.decStore.DecreaseLikedCount(ctx, restaurantId); err != nil {
			log.Println(err)
		}
	}()

	// // Side effect
	// j := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.incStore.IncreaseRestaurant(ctx, data.RestaurantId)
	// })

	// if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	// 	log.Println(err)
	// }

	return nil
}
