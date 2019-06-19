package controllers

import (
	"io/ioutil"
	"net/http"
	"vms-go/goweb/settings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/sagacao/goworld/engine/gwlog"
)

// GetPayList @Get
func GetPayList(c *gin.Context) {
	page := c.Query("page")
	channel := c.Query("channel")
	sdate := c.Query("stime")
	edate := c.Query("etime")
	user := c.DefaultQuery("user", "anonymous")

	gwlog.Debug("GetPayList----------", user, page, channel, sdate, edate)
	urlstr := settings.SvrConfig.Env.URL + "/pay/list/get"
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		gwlog.Error("GetPayList NewRequest err", err)
		return
	}
	query := req.URL.Query()
	query.Add("channel", channel)
	query.Add("sdate", sdate)
	query.Add("edate", edate)
	req.URL.RawQuery = query.Encode()
	gwlog.Debug("GetPayList ", req.URL.String())
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		gwlog.Error("GetPayList DefaultClient Do err", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gwlog.Error("GetPayList ReadAll err", err)
		return
	}
	gwlog.Debug("GetPayList:body", string(body))

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
