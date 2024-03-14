package restaurantbiz

import (
	"200lab-project-1/common"
	restaurantmodel "200lab-project-1/module/restaurant/model"
	"context"

	"gorm.io/gorm"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		ctc context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{
		store: store,
	}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		if err != gorm.ErrInvalidDB {
			return common.ErrEntityNotFound(restaurantmodel.EmtityName, err)
		}
		return common.ErrDB(err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EmtityName, nil)
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EmtityName, err)
	}
	return nil
}
