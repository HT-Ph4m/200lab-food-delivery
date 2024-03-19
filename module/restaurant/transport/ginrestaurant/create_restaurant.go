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

func CreateRestaurant(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMaiDBConnection()

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		// create store
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
