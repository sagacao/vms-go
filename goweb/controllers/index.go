package controllers

import (
	"net/http"
	"strings"
	"vms-go/goweb/dbs"
	"vms-go/goweb/filters/auth"
	"vms-go/goweb/logger"
	"vms-go/goweb/models"
	"vms-go/goweb/settings"
	"vms-go/goweb/utils"

	"github.com/gin-gonic/gin"
)

func IndexGM(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}

// GetUserInfo @Get
func GetUserInfo(c *gin.Context) {
	user := c.DefaultQuery("user", "anonymous")
	logger.Debug("GetUserInfo --------------", user)
	roles, ok := settings.GetUserPrivilege(user)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"role": roles,
		"name": user,
	})
}

// Logout @Post
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// Login @Post
func Login(c *gin.Context) {
	var userCredentials struct {
		User     string `json:"user"  form:"user" binding:"required"`
		Password string `json:"password"  form:"password" binding:"required"`
	}

	// logger.Debug("######################################")
	// user := c.PostForm("user")
	// logger.Debug(user)
	// logger.Debug(c.Request.ParseForm())
	// logger.Debug(c.Request.Body)
	// logger.Debug(c.ContentType())
	// data, _ := ioutil.ReadAll(c.Request.Body)
	// logger.Debug("ctx.Request.body:", string(data))

	if err := c.Bind(&userCredentials); err != nil {
		logger.Error(err)
		return
	}

	pwd, ok := settings.GetUser(userCredentials.User)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户不存在",
		})
		return
	}

	logger.Debug("Login --------------", userCredentials.User, userCredentials.Password, utils.MD5V(pwd))
	if userCredentials.Password != utils.MD5V(pwd) {
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "密码错误",
		})
		return
	}

	authDr, _ := c.MustGet("jwt_auth").(*auth.Auth)
	token, _ := (*authDr).Login(c.Request, c.Writer, map[string]interface{}{userCredentials.User: userCredentials.Password}).(string)

	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"token": token,
	})
}

// GetLogger @Get
func GetLogger(c *gin.Context) {
	page := c.Query("page")
	user := c.DefaultQuery("user", "anonymous")
	stime := c.Query("stime")
	etime := c.Query("etime")

	logger.Debug("Logger----------", user, page, stime, etime)
	//reply := make([]*models.LogStats, 0)
	var reply []*models.LogStats
	dbs.GetDBService().QueryLoggerStats(user, stime, etime, &reply)
	if strings.Contains(page, "cut") {
		var cutReply []*models.LogStats
		for _, v := range reply {
			stats := &models.LogStats{}
			stats.Channel = v.Channel
			stats.Game = v.Game
			stats.LogDate = v.LogDate

			rete := utils.ToFloat64(v.Retention) / 100.00
			stats.Newly = utils.ToString(utils.Round(utils.ToFloat64(v.Newly) * rete))
			stats.TowPr = v.TowPr
			stats.ThreePr = v.ThreePr
			stats.SevenPr = v.SevenPr
			stats.Retention = "0"
			cutReply = append(cutReply, stats)
		}

		c.JSON(http.StatusOK, gin.H{
			"list": cutReply,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"list": reply,
		})
	}
}

// SetLogger @Post
func SetLogger(c *gin.Context) {
	var logStats models.LogStats
	if err := c.Bind(&logStats); err != nil {
		logger.Error(err)
		return
	}

	logger.Debug(logStats)

	errcode := 0
	err := dbs.GetDBService().ReplaceLoggerStats(
		logStats.Channel,
		logStats.Game,
		logStats.Newly,
		logStats.TowPr,
		logStats.ThreePr,
		logStats.SevenPr,
		logStats.Retention,
		logStats.LogDate)
	if err != nil {
		errcode = 1
	}

	c.JSON(http.StatusOK, gin.H{"code": errcode, "msg": logStats.Channel})
}

// DelLogger @Post
func DelLogger(c *gin.Context) {
	var logStats models.LogStats
	if err := c.Bind(&logStats); err != nil {
		logger.Error(err)
		return
	}

	logger.Debug(logStats)

	errcode := 0
	err := dbs.GetDBService().RemoveLoggerStats(
		logStats.Channel,
		logStats.Game,
		logStats.LogDate)
	if err != nil {
		errcode = 1
	}

	c.JSON(http.StatusOK, gin.H{"code": errcode, "msg": err})
}
