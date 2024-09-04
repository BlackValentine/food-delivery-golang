package restaurantstorage

import (
	"context"
	"server/common"
	restaurantmodel "server/module/restaurant/model"

	"gorm.io/gorm"
)

func (s *sqlStore) UpdateRestaurant(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *restaurantmodel.RestaurantUpdate,
) error {
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) IncreaseLikedCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikedCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
