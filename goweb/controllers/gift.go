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

// GetGiftCodes @Get
func GetGiftCodes(c *gin.Context) {
	page := c.Query("page")
	game := c.Query("game")
	user := c.DefaultQuery("user", "anonymous")

	gwlog.Debug("Logger----------", user, page, game)
	urlstr := settings.SvrConfig.Env.URL + "/game/gift/get"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		gwlog.Error("GetGiftCodes NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("game", game)
	req.URL.RawQuery = query.Encode()
	gwlog.Debug("GetGiftCodes ", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		gwlog.Error("GetGiftCodes DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("GetGiftCodes ReadAll err", err)
		return
	}
	gwlog.Debug("GetGiftCodes:body", string(body))

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

// GenGifts @Post
func GenGifts(c *gin.Context) {
	var inmsg struct {
		GameID string `json:"game"  form:"game" binding:"required"`
		Number string `json:"number"  form:"number" binding:"required"`
		AType  string `json:"atype"  form:"atype" binding:"required"`
		ACount string `json:"acount"  form:"acount" binding:"required"`
	}
	if err := c.Bind(&inmsg); err != nil {
		gwlog.Error(err)
		return
	}

	gwlog.Debug(inmsg)
	urlstr := settings.SvrConfig.Env.URL + "/game/gift/gen"
	resp, err := http.PostForm(urlstr,
		url.Values{
			"game":   {inmsg.GameID},
			"number": {inmsg.Number},
			"atype":  {inmsg.AType},
			"acount": {inmsg.ACount},
		})
	if err != nil {
		gwlog.Error("GenGifts PostForm err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("GenGifts ReadAll err", err)
		return
	}
	gwlog.Debug("GenGifts:body", string(body))

	c.JSON(http.StatusOK, gin.H{"data": gin.H{}, "errorCode": 0})
}
