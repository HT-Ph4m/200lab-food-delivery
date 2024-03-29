package restaurantbiz

import (
	"200lab-project-1/common"
	restaurantmodel "200lab-project-1/module/restaurant/model"
	"context"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{
		store: store,
	}
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) (*restaurantmodel.RestaurantCreate, error) {
	// emplement logic here
	if err := data.Validate(); err != nil {
		return data, common.ErrInvalidRequest(err)
	}

	if err := biz.store.Create(context, data); err != nil {
		return data, common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return data, nil
}
