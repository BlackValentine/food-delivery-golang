package userstorage

import (
	"context"
	"server/common"
	usermodel "server/module/user/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(
	ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	var user usermodel.User
	if err := s.db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &user, nil
}
