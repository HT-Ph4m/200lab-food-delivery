package ginuser

import (
	"200lab-project-1/common"
	"200lab-project-1/component/appctx"
	"200lab-project-1/component/hasher"
	"200lab-project-1/module/user/userbiz"
	"200lab-project-1/module/user/usermodel"
	"200lab-project-1/module/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		repo := userbiz.NewRegisterBusiness(store, md5)

		if err := repo.Register(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		data.Mask(false)

		c.JSON(http.StatusOK, data.FakeId.String())
	}
}
