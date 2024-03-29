package restaurantstorage

import (
	"200lab-project-1/common"
	restaurantmodel "200lab-project-1/module/restaurant/model"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(
		restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
