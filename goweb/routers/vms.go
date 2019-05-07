package routers

import (
	"vms-go/goweb/controllers"
	"vms-go/goweb/filters/auth"

	"github.com/gin-gonic/gin"
)

func registerRouter(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)

	vms := router.Group("/vms")
	vms.Use(auth.Middleware("jwt"))
	{
		vms.GET("/index", controllers.IndexGM)
		vms.GET("/user/info", controllers.GetUserInfo)
		vms.GET("/logger", controllers.GetLogger)
		vms.POST("/logger", controllers.SetLogger)

		// gm.GET("/notice/search", controllers.GMNoticeSearch)
		// gm.POST("/notice/add", controllers.GMNoticeAdd)
		// gm.POST("/notice/del", controllers.GMNoticeDel)

		// gm.GET("/funcswitch/search", controllers.GMFuncSwitchSearch)
		// gm.POST("/funcswitch/add", controllers.GMFuncSwitchAdd)
		// gm.POST("/funcswitch/del", controllers.GMFuncSwitchDel)

		// gm.GET("/magic/search", controllers.GMMagicSearch)
		// gm.POST("/magic/add", controllers.GMMagicAdd)
		// gm.POST("/magic/del", controllers.GMMagicDel)

		// gm.GET("/gates/search", controllers.GMGatesSearch)
		// gm.POST("/gates/del", controllers.GMGatesDel)

		// gm.GET("/config/search", controllers.GMServerConfigSearch)
		// gm.POST("/config/update", controllers.GMSetServerConfig)
	}
}
