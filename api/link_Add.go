package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/link/handler"
	"github.com/ismdeep/link/request"
	"github.com/ismdeep/log"
)

// LinkAdd add new link
// @Summary add new link
// @Author l.jiang.1024@gmail.com
// @Description add new link
// @Tags Link
// @Router	/api/v1/links [POST]
func LinkAdd(c *gin.Context) {
	req := &request.LinkSave{}
	if err := c.ShouldBindJSON(req); err != nil {
		log.Error("LinkAdd", "req", req, "err", err)
		Fail(c, err.Error())
		return
	}

	respData, err := handler.Link.Add(req)
	if err != nil {
		log.Error("LinkAdd", "req", req, "err", err)
		Fail(c, err.Error())
		return
	}

	Success(c, "", respData)
}
