package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/link/config"
	"github.com/ismdeep/link/model"
	"github.com/ismdeep/log"
)

// LinkJump jump link
// @Summary jump link
// @Author l.jiang.1024@gmail.com
// @Description jump link
// @Tags Link
// @Router	/:link_id [get]
func LinkJump(c *gin.Context) {
	linkID := c.Param("link_id")
	log.Info("LinkJump", "link_id", linkID)
	link := &model.Link{}
	if err := model.DB.Where("id=?", linkID).First(link).Error; err != nil {
		c.Redirect(302, fmt.Sprintf("%v/404", config.Config.Server.Site))
		return
	}

	c.Redirect(301, link.Target)
	return
}
