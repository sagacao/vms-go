package auth

import (
	"log"
	"net/http"
	"vms-go/goweb/filters/auth/drivers"

	"github.com/gin-gonic/gin"
)

var driverList = map[string]func() Auth{
	"cookie": func() Auth {
		return drivers.NewCookieAuthDriver()
	},
	"jwt": func() Auth {
		return drivers.NewJwtAuthDriver()
	},
}

type Auth interface {
	Check(c *gin.Context) bool
	User(c *gin.Context) interface{}
	Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{}
	Logout(http *http.Request, w http.ResponseWriter) bool
}

func RegisterGlobalAuthDriver(authKey string, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		driver := GenerateAuthDriver(authKey)
		c.Set(key, driver)
		c.Next()
	}
}

func Middleware(authKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		driver := GenerateAuthDriver(authKey)
		if !(*driver).Check(c) {
			//c.Redirect(http.StatusUnauthorized, "../login")
			// c.HTML(http.StatusOK, "login.html", gin.H{
			// 	"title": "尚未登录，请登录",
			// })
			log.Printf("ttttttttttttttttttttt")
			c.AbortWithStatus(http.StatusUnauthorized)
			//c.Abort()
		}
		c.Next()
	}
}

func GenerateAuthDriver(string string) *Auth {
	var authDriver Auth
	authDriver = driverList[string]()
	return &authDriver
}

func GetCurUser(c *gin.Context, key string) map[string]interface{} {
	authDriver, _ := c.MustGet(key).(*Auth)
	return (*authDriver).User(c).(map[string]interface{})
}

func User(c *gin.Context) map[string]interface{} {
	return GetCurUser(c, "jwt_auth")
}
