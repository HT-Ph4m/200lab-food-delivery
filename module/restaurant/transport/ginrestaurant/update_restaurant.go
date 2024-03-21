package ginrestaurant

import (
	"200lab-project-1/common"
	"200lab-project-1/component/appctx"
	restaurantbiz "200lab-project-1/module/restaurant/biz"
	restaurantmodel "200lab-project-1/module/restaurant/model"
	restaurantstorage "200lab-project-1/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMaiDBConnection()

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)

		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		restaurant, err := biz.UpdateRestaurant(c.Request.Context(), int(uid.GetLocalID()), &data)

		if err != nil {
			panic(err)
		}

		restaurant.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(restaurant))

	}
}
