package ginuser

import (
	"200lab-project-1/common"
	"200lab-project-1/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, data)
	}
}
