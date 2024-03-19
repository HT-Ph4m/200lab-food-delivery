package restaurantbiz

import (
	"200lab-project-1/common"
	restaurantmodel "200lab-project-1/module/restaurant/model"
	"context"

	"gorm.io/gorm"
)

type FindRestaurantStore interface {
	FindDataWithCondition(
		ctc context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)
}

type findRestaurantBiz struct {
	store FindRestaurantStore
}

func NewFinRestaurantBiz(store FindRestaurantStore) *findRestaurantBiz {
	return &findRestaurantBiz{
		store: store,
	}
}

func (biz *findRestaurantBiz) FindRestaurant(ctx context.Context, id string) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err != gorm.ErrInvalidDB {
			return nil, common.ErrEntityNotFound(restaurantmodel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}

	return result, nil
}
