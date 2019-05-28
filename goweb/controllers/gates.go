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

// ReqSvrGates @Get
func ReqSvrGates(c *gin.Context) {
	game, ok := c.GetQuery("game")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{"errorCode": http.StatusNotAcceptable})
		return
	}
	user := c.DefaultQuery("user", "anonymous")
	gwlog.Debug("ReqSvrGates --------------", user, game)

	urlstr := settings.SvrConfig.Env.URL + "/gates/getGates"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		gwlog.Error("ReqSvrGates NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("gameId", game)
	req.URL.RawQuery = query.Encode()
	gwlog.Debug("ReqSvrGates ", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		gwlog.Error("ReqSvrGates DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("ReqSvrCfg ReadAll err", err)
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

// EditSvrGates @Post
func EditSvrGates(c *gin.Context) {
	var editData struct {
		Game  string `json:"game"  form:"game" binding:"required"`
		Name  string `json:"name"  form:"name" binding:"required"`
		Value string `json:"value"  form:"value" binding:"required"`
		Svr   string `json:"svr"  form:"svr" binding:"required"`
	}

	if err := c.Bind(&editData); err != nil {
		gwlog.Error(err)
		return
	}
	gwlog.Info("EditSvrGates", editData.Game, editData.Value)

	urlstr := settings.SvrConfig.Env.URL + "/Logic/user/setServerGates"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"gameId": {editData.Game},
			"name":   {editData.Name},
			"value":  {editData.Value},
			"count":  {"0"},
			"svr":    {editData.Svr},
		})
	if err != nil {
		gwlog.Error("EditSvrGates PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("EditSvrGates ReadAll err", err)
		return
	}
	gwlog.Debug("EditSvrGates:body", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}

// RemoveSvrGates @Post
func RemoveSvrGates(c *gin.Context) {
	var editData struct {
		Game string `json:"game"  form:"game" binding:"required"`
		Name string `json:"name"  form:"name" binding:"required"`
	}

	if err := c.Bind(&editData); err != nil {
		gwlog.Error(err)
		return
	}
	gwlog.Info("RemoveSvrGates", editData.Game, editData.Name)

	urlstr := settings.SvrConfig.Env.URL + "/gates/delGates"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"gameId": {editData.Game},
			"name":   {editData.Name},
		})
	if err != nil {
		gwlog.Error("RemoveSvrGates PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("RemoveSvrGates ReadAll err", err)
		return
	}
	gwlog.Debug("RemoveSvrGates:body", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}
