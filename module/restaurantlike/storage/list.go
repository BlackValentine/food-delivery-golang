package restaurantlikestorage

import (
	"context"
	"fmt"
	"server/common"
	restaurantlikemodel "server/module/restaurantlike/model"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) GetRestautrantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id;"`
		LikeCount    int `gorm:"column:count;"`
	}

	var listLike []sqlData

	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil
}

func (s *sqlStore) GetUsersLikeRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var result []restaurantlikemodel.Like

	db := s.db

	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02T15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("created_at DESC").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User

		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil
}
