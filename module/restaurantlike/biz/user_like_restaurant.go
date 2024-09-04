package restaurantlikebiz

import (
	"context"
	"log"
	restaurantlikemodel "server/module/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncLikedCountRestaurantStore interface {
	IncreaseLikedCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncLikedCountRestaurantStore
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	incStore IncLikedCountRestaurantStore,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:    store,
		incStore: incStore,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.CreateRestaurantLike(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	go func() {
		if err := biz.incStore.IncreaseLikedCount(ctx, data.RestaurantId); err != nil {
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
