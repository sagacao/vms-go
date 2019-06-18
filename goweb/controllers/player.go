package controllers

import (
	"net/http"
	"nuvem/engine/coder"
	"vms-go/goweb/dbs"

	"github.com/gin-gonic/gin"
	"github.com/sagacao/goworld/engine/gwlog"
)

// ReqPlayerInfo @Get
func ReqPlayerInfo(c *gin.Context) {
	game, ok := c.GetQuery("game")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{"errorCode": http.StatusNotAcceptable})
		return
	}
	uid, ok := c.GetQuery("uid")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{"errorCode": http.StatusNotAcceptable})
		return
	}
	user := c.DefaultQuery("user", "anonymous")
	gwlog.Debug("ReqPlayerInfo --------------", user, game)

	var reply coder.JSON
	err := dbs.GetDBService().QueryPlayerInfoByUid("heros", uid, reply)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": reply, "errorCode": 9100})
	}

	c.JSON(http.StatusOK, reply)
}

// EditPlayerInfo @Post
func EditPlayerInfo(c *gin.Context) {
	var editData struct {
		Game string `json:"game"  form:"game" binding:"required"`
		UID  string `json:"uid"  form:"uid" binding:"required"`
		Data string `json:"data"  form:"data" binding:"required"`
	}

	if err := c.Bind(&editData); err != nil {
		gwlog.Error(err)
		return
	}
	gwlog.Info("EditPlayerInfo", editData.Game, editData.UID)

	var jsondata coder.JSON
	err := coder.ToJSON([]byte(editData.Data), jsondata)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": err, "errorCode": 9999})
		return
	}
	gwlog.Debug(jsondata)

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}
