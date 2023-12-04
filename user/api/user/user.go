package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerUser struct {
}

func (*HandlerUser) getCapture(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hello"})
}
