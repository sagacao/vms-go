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

// ReqSvrCfg @Get
func ReqSvrCfg(c *gin.Context) {
	game, ok := c.GetQuery("game")
	if !ok {
		c.JSON(http.StatusNotAcceptable, gin.H{"errorCode": http.StatusNotAcceptable})
		return
	}
	user := c.DefaultQuery("user", "anonymous")
	gwlog.Debugf("ReqSvrCfg --------------%s %s ", user, game)

	urlstr := settings.SvrConfig.Env.URL + "/Logic/user/getServerConfig"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		gwlog.Error("ReqSvrCfg NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("gameId", game)
	req.URL.RawQuery = query.Encode()
	gwlog.Debugf("ReqSvrCfg %s", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		gwlog.Error("ReqSvrCfg DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("ReqSvrCfg ReadAll err", err)
		return
	}
	//gwlog.Debugf("ReqSvrCfg:body %s", string(body))

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

// EditSvrCfg @Post
func EditSvrCfg(c *gin.Context) {
	var editData struct {
		Game  string `json:"game"  form:"game" binding:"required"`
		Value string `json:"value"  form:"value" binding:"required"`
	}

	if err := c.Bind(&editData); err != nil {
		gwlog.Error(err)
		return
	}
	gwlog.Infof("EditSvrCfg %s %s", editData.Game, editData.Value)

	urlstr := settings.SvrConfig.Env.URL + "/Logic/user/setServerConfig"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		gwlog.Error("EditSvrCfg NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("gameId", editData.Game)
	query.Add("value", editData.Value)
	req.URL.RawQuery = query.Encode()
	gwlog.Debugf("EditSvrCfg %s", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		gwlog.Error("EditSvrCfg DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("EditSvrCfg ReadAll err", err)
		return
	}
	gwlog.Debugf("EditSvrCfg:body %s", string(body))

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
	gwlog.Debugf("ReqSvrSwitch ------------- %s %s", user, game)

	urlstr := settings.SvrConfig.Env.URL + "/forbiden/getForbiden"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		gwlog.Error("ReqSvrSwitch NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("gameId", game)
	req.URL.RawQuery = query.Encode()
	gwlog.Debugf("ReqSvrSwitch %s", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		gwlog.Error("ReqSvrSwitch DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("ReqSvrSwitch ReadAll err", err)
		return
	}
	gwlog.Debugf("ReqSvrSwitch:body %s", string(body))

	args := make(gin.H)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(body, &args)
	if err != nil {
		gwlog.Error("json.Unmarshal", err)
		c.JSON(http.StatusOK, args)
	} else {
		//utils.DumpSocketData(args["data"])
		rspdata, ok := args["data"].(map[string]interface{})
		if !ok {
			gwlog.Error("json. no data")
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
		gwlog.Error(err)
		return
	}
	gwlog.Infof("EditSvrSwitch %s %s %s", editData.Game, editData.Funcname, editData.Funcswitch)

	urlstr := settings.SvrConfig.Env.URL + "/forbiden/setForbiden"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"gameId": {editData.Game},
			"name":   {editData.Funcname},
			"flag":   {editData.Funcswitch},
		})
	if err != nil {
		gwlog.Error("EditSvrSwitch PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("EditSvrSwitch ReadAll err", err)
		return
	}
	gwlog.Debugf("EditSvrSwitch:body %s", string(body))

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
		gwlog.Error(err)
		return
	}
	gwlog.Infof("RemoveSvrSwitch %s %s %s ", editData.Game, editData.Funcname, editData.Funcswitch)

	urlstr := settings.SvrConfig.Env.URL + "/forbiden/delForbiden"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"gameId": {editData.Game},
			"name":   {editData.Funcname},
		})
	if err != nil {
		gwlog.Error("RemoveSvrSwitch PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("RemoveSvrSwitch ReadAll err", err)
		return
	}
	gwlog.Debugf("RemoveSvrSwitch:body %s", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}
