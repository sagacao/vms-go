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
		vms.GET("/stats/get", controllers.GetLogger)
		vms.POST("/stats/edit", controllers.SetLogger)
		vms.POST("/stats/rm", controllers.DelLogger)

		vms.GET("/cfg/req", controllers.ReqSvrCfg)
		vms.POST("/cfg/edit", controllers.EditSvrCfg)

		vms.GET("/switch/req", controllers.ReqSvrSwitch)
		vms.POST("/switch/edit", controllers.EditSvrSwitch)
		vms.POST("/switch/remove", controllers.RemoveSvrSwitch)

		// gm.GET("/notice/search", controllers.GMNoticeSearch)
		// gm.POST("/notice/add", controllers.GMNoticeAdd)
		// gm.POST("/notice/del", controllers.GMNoticeDel)

		// gm.GET("/magic/search", controllers.GMMagicSearch)
		// gm.POST("/magic/add", controllers.GMMagicAdd)
		// gm.POST("/magic/del", controllers.GMMagicDel)

		vms.GET("/gates/req", controllers.ReqSvrGates)
		vms.POST("/gates/edit", controllers.EditSvrGates)
		vms.POST("/gates/remove", controllers.RemoveSvrGates)
	}
}
