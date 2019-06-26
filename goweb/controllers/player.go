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

var (
	gamemap = map[string]string{
		"21001": "heros",
	}
)

// ReqPlayerInfo @Get
func ReqPlayerInfo(c *gin.Context) {
	game, ok := c.GetQuery("game")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{"errorCode": http.StatusNotAcceptable})
		return
	}
	key, ok := c.GetQuery("key")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{"errorCode": http.StatusNotAcceptable})
		return
	}
	ktype, ok := c.GetQuery("ktype")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{"errorCode": http.StatusNotAcceptable})
		return
	}

	user := c.DefaultQuery("user", "anonymous")
	gwlog.Debug("ReqPlayerInfo --------------", user, game, key, ktype)
	urlstr := settings.SvrConfig.Env.URL + "/game/user/get"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		gwlog.Error("ReqPlayerInfo NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("game", gamemap[game])
	query.Add("uid", key)
	query.Add("type", ktype)
	req.URL.RawQuery = query.Encode()
	gwlog.Debug("ReqPlayerInfo ", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		gwlog.Error("ReqPlayerInfo DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("ReqPlayerInfo ReadAll err", err)
		return
	}
	gwlog.Debug("ReqPlayerInfo:body", string(body))

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

// EditPlayerInfo @Post
func EditPlayerInfo(c *gin.Context) {
	var inmsg struct {
		Game string `json:"game"  form:"game" binding:"required"`
		UID  string `json:"uid"  form:"uid" binding:"required"`
		Data string `json:"data"  form:"data" binding:"required"`
	}

	if err := c.Bind(&inmsg); err != nil {
		gwlog.Error(err)
		return
	}
	gwlog.Info("EditPlayerInfo", inmsg.Game, inmsg.UID)

	gwlog.Debug(inmsg)
	urlstr := settings.SvrConfig.Env.URL + "/game/user/edit"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"game": {gamemap[inmsg.Game]},
			"uid":  {inmsg.UID},
			"data": {inmsg.Data},
		})
	if err != nil {
		gwlog.Error("EditPlayerInfo PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("EditPlayerInfo ReadAll err", err)
		return
	}
	gwlog.Debug("EditPlayerInfo:body", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}
