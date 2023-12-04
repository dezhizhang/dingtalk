package user

import (
	common "digit.icu.common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerUser struct {
}

func (*HandlerUser) getCapture(c *gin.Context) {
	rsp := &common.Result{}
	c.JSON(http.StatusOK, rsp.Success("12356"))
}
