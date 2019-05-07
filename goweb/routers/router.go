package routers

import (
	"net/http"
	"vms-go/goweb/filters"
	"vms-go/goweb/filters/auth"
	"vms-go/goweb/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("dist/*.html")        // 添加入口index.html
	router.LoadHTMLFiles("static/*/*")        // 添加资源路径
	router.Static("/static", "./dist/static") // 添加资源路径
	router.StaticFile("/", "dist/index.html") //前端接口

	router.Use(gin.Logger())
	router.Use(handleErrors())            // 错误处理
	router.Use(filters.RegisterSession()) // 全局session
	router.Use(filters.RegisterCache())   // 全局cache

	router.Use(auth.RegisterGlobalAuthDriver("cookie", "web_auth")) // 全局auth cookie
	router.Use(auth.RegisterGlobalAuthDriver("jwt", "jwt_auth"))    // 全局auth jwt

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "no router",
		})
		return
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "no method",
		})
		return
	})

	registerRouter(router)
	return router
}

func handleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				logger.Error(err)

				var (
					errMsg     string
					mysqlError *mysql.MySQLError
					ok         bool
				)
				if errMsg, ok = err.(string); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg":  "system error, " + errMsg,
					})
					return
				} else if mysqlError, ok = err.(*mysql.MySQLError); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg":  "system error, " + mysqlError.Error(),
					})
					return
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg":  "system error",
					})
					return
				}
			}
		}()
		c.Next()
	}
}
