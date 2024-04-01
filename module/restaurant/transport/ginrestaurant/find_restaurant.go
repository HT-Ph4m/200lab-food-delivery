package ginrestaurant

import (
	"200lab-project-1/common"
	"200lab-project-1/component/appctx"
	restaurantbiz "200lab-project-1/module/restaurant/biz"
	restaurantstorage "200lab-project-1/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindDataWithCondition(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)

		biz := restaurantbiz.NewFindRestaurantBiz(store)

		result, err := biz.FindRestaurant(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		result.Mask(false)

		c.JSON(http.StatusOK, result)
	}
}
