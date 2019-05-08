package controllers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"vms-go/goweb/logger"
	"vms-go/goweb/settings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

// ReqSvrCfg @Get
func ReqSvrCfg(c *gin.Context) {
	game, ok := c.GetQuery("game")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{"errorCode": http.StatusNotAcceptable})
		return
	}
	user := c.DefaultQuery("user", "anonymous")
	logger.Debug("ReqSvrCfg --------------", user, game)

	urlstr := settings.SvrConfig.Env.URL + "/Logic/user/getServerConfig"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		logger.Error("ReqSvrCfg NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("gameId", game)
	req.URL.RawQuery = query.Encode()
	logger.Debug("ReqSvrCfg ", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("ReqSvrCfg DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("ReqSvrCfg ReadAll err", err)
		return
	}
	//logger.Debug("ReqSvrCfg:body", string(body))

	args := make(gin.H)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(body, &args)
	if err != nil {
		logger.Error("json.Unmarshal", err)
		c.JSON(http.StatusOK, args)
	} else {
		c.JSON(http.StatusOK, args)
	}
}

// EditSvrCfg @Post
func EditSvrCfg(c *gin.Context) {
	var editData struct {
		Game  string `json:"game"  form:"game" binding:"required"`
		Value string `json:"value"  form:"value" binding:"required"`
	}

	if err := c.Bind(&editData); err != nil {
		logger.Error(err)
		return
	}
	logger.Info("EditSvrCfg", editData.Game, editData.Value)

	urlstr := settings.SvrConfig.Env.URL + "/Logic/user/setServerConfig"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		logger.Error("EditSvrCfg NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("gameId", editData.Game)
	query.Add("value", editData.Value)
	req.URL.RawQuery = query.Encode()
	logger.Debug("EditSvrCfg ", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("EditSvrCfg DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("EditSvrCfg ReadAll err", err)
		return
	}
	logger.Debug("EditSvrCfg:body", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}

// ReqSvrSwitch @Get
func ReqSvrSwitch(c *gin.Context) {
	game, ok := c.GetQuery("game")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{"errorCode": http.StatusNotAcceptable})
		return
	}
	user := c.DefaultQuery("user", "anonymous")
	logger.Debug("ReqSvrSwitch --------------", user, game)

	urlstr := settings.SvrConfig.Env.URL + "/forbiden/getForbiden"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		logger.Error("ReqSvrSwitch NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("gameId", game)
	req.URL.RawQuery = query.Encode()
	logger.Debug("ReqSvrSwitch ", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("ReqSvrSwitch DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("ReqSvrSwitch ReadAll err", err)
		return
	}
	logger.Debug("ReqSvrSwitch:body", string(body))

	args := make(gin.H)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(body, &args)
	if err != nil {
		logger.Error("json.Unmarshal", err)
		c.JSON(http.StatusOK, args)
	} else {
		//utils.DumpSocketData(args["data"])
		rspdata, ok := args["data"].(map[string]interface{})
		if !ok {
			logger.Error("json. no data")
			c.JSON(http.StatusOK, args)
		} else {
			var rsp []gin.H
			for k, v := range rspdata {
				rsp = append(rsp, gin.H{"funcname": k, "funcswitch": v})
			}
			c.JSON(http.StatusOK, gin.H{"data": rsp, "errorCode": 0})
		}
	}
}

// EditSvrSwitch @Post
func EditSvrSwitch(c *gin.Context) {
	var editData struct {
		Game       string `json:"game"  form:"game" binding:"required"`
		Funcname   string `json:"funcname"  form:"funcname" binding:"required"`
		Funcswitch string `json:"funcswitch"  form:"funcswitch" `
	}

	if err := c.Bind(&editData); err != nil {
		logger.Error(err)
		return
	}
	logger.Info("EditSvrSwitch", editData.Game, editData.Funcname, editData.Funcswitch)

	urlstr := settings.SvrConfig.Env.URL + "/forbiden/setForbiden"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"gameId": {editData.Game},
			"name":   {editData.Funcname},
			"flag":   {editData.Funcswitch},
		})
	if err != nil {
		logger.Error("EditSvrSwitch PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("EditSvrSwitch ReadAll err", err)
		return
	}
	logger.Debug("EditSvrSwitch:body", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}

// RemoveSvrSwitch @Post
func RemoveSvrSwitch(c *gin.Context) {
	var editData struct {
		Game       string `json:"game"  form:"game" binding:"required"`
		Funcname   string `json:"funcname"  form:"funcname" binding:"required"`
		Funcswitch string `json:"funcswitch"  form:"funcswitch" `
	}

	if err := c.Bind(&editData); err != nil {
		logger.Error(err)
		return
	}
	logger.Info("RemoveSvrSwitch", editData.Game, editData.Funcname, editData.Funcswitch)

	urlstr := settings.SvrConfig.Env.URL + "/forbiden/delForbiden"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"gameId": {editData.Game},
			"name":   {editData.Funcname},
		})
	if err != nil {
		logger.Error("RemoveSvrSwitch PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("RemoveSvrSwitch ReadAll err", err)
		return
	}
	logger.Debug("RemoveSvrSwitch:body", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}
