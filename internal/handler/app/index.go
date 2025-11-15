package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndex(c *gin.Context) {
	c.Redirect(http.StatusFound, "https://yunwanjia.me")
}
