package routers

import (
	"demo/controller"
	"demo/middleware"
	"demo/util"

	"github.com/gin-gonic/gin"
)

// func defaultFunc(c *gin.Context) {
// 	c.String(http.StatusOK, "默认处理函数")
// }
func InitRouters() {
	// gin.SetMode(gin.ReleaseMode) //发布版本，未设置默认为DEBUG
	r := gin.Default()
	// r.NoRoute()	404页面
	v1 := r.Group("api")
	{
		v1.POST("/register", controller.SignUp)
		v1.POST("/register/sendMessage", controller.SignUpSendMessage)
		v1.POST("/login", controller.SignIn)
		v1.POST("/findpassword", controller.FindPassword)
		v1.GET("/testAu", middleware.JWTAuth(), func(c *gin.Context) {
			c.JSON(200, gin.H{
				"Status":  200,
				"Message": "test au ok!",
			})
		})
	}

	port := util.GetConfig("port")
	r.Run(port)
}
