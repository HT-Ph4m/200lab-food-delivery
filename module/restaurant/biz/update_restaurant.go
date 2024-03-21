package restaurantbiz

import (
	"200lab-project-1/common"
	restaurantmodel "200lab-project-1/module/restaurant/model"
	"context"

	"gorm.io/gorm"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(
		ctc context.Context,
		condition map[string]interface{},
		moreKey ...string) (*restaurantmodel.Restaurant, error)
	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{
		store: store,
	}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(context context.Context, id int, data *restaurantmodel.RestaurantUpdate) (*restaurantmodel.Restaurant, error) {
	// Tìm kiếm dữ liệu nhà hàng với ID cung cấp
	restaurant, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		// Kiểm tra xem lỗi có phải là lỗi cơ sở dữ liệu không hợp lệ không
		if err == gorm.ErrInvalidDB {
			return nil, common.ErrDB(err)
		}
		// Nếu không tìm thấy nhà hàng, trả về lỗi "không tìm thấy thực thể"
		return nil, common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	// Thực hiện cập nhật dữ liệu của nhà hàng
	if err := biz.store.Update(context, id, data); err != nil {
		return nil, common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	// Cập nhật bản ghi nhà hàng với dữ liệu mới
	restaurant.Name = *data.Name
	restaurant.Addr = *data.Addr
	// Cập nhật các trường khác nếu có

	// Trả về bản ghi nhà hàng sau khi cập nhật và không có lỗi
	return restaurant, nil
}
