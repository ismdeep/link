package api

import "github.com/gin-gonic/gin"

// EchoHelp echo help
// @Summary echo help
// @Author l.jiang.1024@gmail.com
// @Description echo help
// @Tags Link
// @Router	/ [GET]
func EchoHelp(c *gin.Context) {
	_, _ = c.Writer.WriteString(`Welcome to Link.

GET  /             => Home Page.
GET  /404          => Not Found.
POST /api/v1/links => Add link
GET  /:link_id     => Jump to link

--------------------------------------------

{
    "id": "",
    "target": "",
    "description": ""
}

--------------------------------------------

`)
}
