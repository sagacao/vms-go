package controllers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"vms-go/goweb/settings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/sagacao/goworld/engine/gwlog"
)

// ReqSvrChannelSatatus @Get
func ReqSvrChannelSatatus(c *gin.Context) {
	game, ok := c.GetQuery("game")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{"errorCode": http.StatusNotAcceptable})
		return
	}
	user := c.DefaultQuery("user", "anonymous")
	gwlog.Debug("ReqSvrChannelSatatus --------------", user, game)

	urlstr := settings.SvrConfig.Env.URL + "/status/getChannelSwitch"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		gwlog.Error("ReqSvrChannelSatatus NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("gameId", game)
	req.URL.RawQuery = query.Encode()
	gwlog.Debug("ReqSvrChannelSatatus ", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		gwlog.Error("ReqSvrChannelSatatus DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("ReqSvrChannelSatatus ReadAll err", err)
		return
	}
	// gwlog.Debug("ReqSvrChannelSatatus:body", string(body))

	args := make(gin.H)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(body, &args)
	if err != nil {
		gwlog.Error("json.Unmarshal", err)
		c.JSON(http.StatusOK, args)
	} else {
		// gwlog.Debug(args)
		c.JSON(http.StatusOK, args)
	}
}

// EditSvrChannelStatus @Post
func EditSvrChannelStatus(c *gin.Context) {
	var editData struct {
		Game    string `json:"game"  form:"game" binding:"required"`
		Channel string `json:"channel"  form:"channel" binding:"required"`
		Name    string `json:"name"  form:"name" binding:"required"`
		Status  string `json:"status"  form:"status" binding:"required"`
		Action  string `json:"action"  form:"action" binding:"required"`
	}

	if err := c.Bind(&editData); err != nil {
		gwlog.Error(err)
		return
	}
	gwlog.Info("EditSvrChannelStatus", editData.Game, editData.Channel)

	urlstr := settings.SvrConfig.Env.URL + "/status/setChannelSwitch"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"gameId":  {editData.Game},
			"channel": {editData.Channel},
			"name":    {editData.Name},
			"status":  {editData.Status},
			"action":  {editData.Action},
		})
	if err != nil {
		gwlog.Error("EditSvrChannelStatus PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("EditSvrChannelStatus ReadAll err", err)
		return
	}
	gwlog.Debug("EditSvrChannelStatus:body", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}

// RemoveSvrChannelStatus @Post
func RemoveSvrChannelStatus(c *gin.Context) {
	var editData struct {
		Game    string `json:"game"  form:"game" binding:"required"`
		Channel string `json:"channel"  form:"channel" binding:"required"`
		Name    string `json:"name"  form:"name" binding:"required"`
	}

	if err := c.Bind(&editData); err != nil {
		gwlog.Error(err)
		return
	}
	gwlog.Info("RemoveSvrChannelStatus", editData.Game, editData.Channel)

	urlstr := settings.SvrConfig.Env.URL + "/status/delChannelSwitch"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"gameId":  {editData.Game},
			"channel": {editData.Channel},
			"name":    {editData.Name},
		})
	if err != nil {
		gwlog.Error("RemoveSvrChannelStatus PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("RemoveSvrChannelStatus ReadAll err", err)
		return
	}
	gwlog.Debug("RemoveSvrChannelStatus:body", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}
