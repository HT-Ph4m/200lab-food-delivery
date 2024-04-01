package ginuser

import (
	"200lab-project-1/common"
	"200lab-project-1/component/appctx"
	"200lab-project-1/component/hasher"
	"200lab-project-1/component/tokenprovider/jwt"
	"200lab-project-1/module/user/userbiz"
	"200lab-project-1/module/user/usermodel"
	"200lab-project-1/module/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey()) //appctx.SecretKey()

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBusiness(appCtx, store, 60*60*24*30, tokenProvider, md5)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, account)
	}
}
