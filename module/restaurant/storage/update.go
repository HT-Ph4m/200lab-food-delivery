package restaurantstorage

import (
	"200lab-project-1/common"
	restaurantmodel "200lab-project-1/module/restaurant/model"
	"context"
)

func (s *sqlStore) Update(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
