package restaurantlikestorage

import (
	"context"
	"server/common"
	restaurantlikemodel "server/module/restaurantlike/model"
)

func (s *sqlStore) CreateRestaurantLike(ctx context.Context, data *restaurantlikemodel.Like) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
