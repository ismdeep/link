package api

import "github.com/gin-gonic/gin"

// EchoNotFound echo not found
// @Summary echo not found
// @Author l.jiang.1024@gmail.com
// @Description echo not found
// @Tags Link
// @Router	/404 [GET]
func EchoNotFound(c *gin.Context) {
	_, _ = c.Writer.WriteString("404 page not found")
}
