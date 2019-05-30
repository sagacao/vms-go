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

// ReqSvrAppInfo @Get
func ReqSvrAppInfo(c *gin.Context) {
	game, ok := c.GetQuery("game")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{"errorCode": http.StatusNotAcceptable})
		return
	}
	user := c.DefaultQuery("user", "anonymous")
	gwlog.Debug("ReqSvrAppInfo --------------", user, game)

	urlstr := settings.SvrConfig.Env.URL + "/vapp/getvapp"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		gwlog.Error("ReqSvrAppInfo NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("gameId", game)
	req.URL.RawQuery = query.Encode()
	gwlog.Debug("ReqSvrAppInfo ", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		gwlog.Error("ReqSvrAppInfo DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("ReqSvrAppInfo ReadAll err", err)
		return
	}
	//gwlog.Debug("ReqSvrCfg:body", string(body))

	args := make(gin.H)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(body, &args)
	if err != nil {
		gwlog.Error("json.Unmarshal", err)
		c.JSON(http.StatusOK, args)
	} else {
		c.JSON(http.StatusOK, args)
	}
}

// EditSvrAppInfo @Post
func EditSvrAppInfo(c *gin.Context) {
	var editData struct {
		Game      string `json:"game"  form:"game" binding:"required"`
		Channel   string `json:"channel"  form:"channel" binding:"required"`
		AppID     string `json:"appid"  form:"appid" binding:"required"`
		AppSecret string `json:"secret"  form:"secret" binding:"required"`
		Desc      string `json:"desc"  form:"desc" binding:"required"`
	}

	if err := c.Bind(&editData); err != nil {
		gwlog.Error(err)
		return
	}
	gwlog.Info("EditSvrAppInfo", editData.Game, editData.Channel)

	urlstr := settings.SvrConfig.Env.URL + "/vapp/setvapp"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"gameId":    {editData.Game},
			"channel":   {editData.Channel},
			"appid":     {editData.AppID},
			"appsecret": {editData.AppSecret},
			"desc":      {editData.Desc},
		})
	if err != nil {
		gwlog.Error("EditSvrAppInfo PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("EditSvrAppInfo ReadAll err", err)
		return
	}
	gwlog.Debug("EditSvrAppInfo:body", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}

// RemoveSvrAppInfo @Post
func RemoveSvrAppInfo(c *gin.Context) {
	var editData struct {
		Game  string `json:"game"  form:"game" binding:"required"`
		AppID string `json:"appid"  form:"appid" binding:"required"`
	}

	if err := c.Bind(&editData); err != nil {
		gwlog.Error(err)
		return
	}
	gwlog.Info("RemoveSvrAppInfo", editData.Game, editData.AppID)

	urlstr := settings.SvrConfig.Env.URL + "/vapp/delvapp"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"gameId": {editData.Game},
			"appid":  {editData.AppID},
		})
	if err != nil {
		gwlog.Error("RemoveSvrAppInfo PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("RemoveSvrAppInfo ReadAll err", err)
		return
	}
	gwlog.Debug("RemoveSvrAppInfo:body", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}
