package router

import (
	"net/http"
	"strings"

	"PersonalizedRecommendationSystem/handler/goods"

	"PersonalizedRecommendationSystem/handler/login"

	"PersonalizedRecommendationSystem/handler/order"

	"PersonalizedRecommendationSystem/handler/profit"

	"PersonalizedRecommendationSystem/handler/signup"

	"PersonalizedRecommendationSystem/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Router *gin.Engine

func init() {
	Router = gin.New()
	Router.MaxMultipartMemory = 200 << 20 // 64 MiB

}

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	//  注册 API 路由
	RegisterAPIRoutes(router)

	//  配置 404 路由
	setup404Handler(router)
}

// 注册全局中间件
func registerGlobalMiddleWare(g *gin.Engine) {
	g.Use(middlewares.Logger())
	g.Use(Cors())
	//g.Use(middlewares.Cors()) //开启中间件 允许使用跨域请求
}

// 设置所有地域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// 注册 API 路由
func RegisterAPIRoutes(g *gin.Engine) {
	g.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g1 := g.Group("/api/v1/login")
	{
		//使用手机号和密码登陆
		g1.POST("/common", login.CommonUserLoginByPhone)
		g1.POST("/staff", login.StaffUserLoginByPhone)
		g1.POST("/member", login.MemberUserLoginByPhone)
	}
	g2 := g.Group("/api/v1/sign")
	{
		//使用手机号和密码注册
		g2.POST("/common", signup.CommonUserSignupUsingPhone)
		g2.POST("/staff", signup.StaffUserSignupUsingPhone)
		g2.POST("/member", signup.MemberUserSignupUsingPhone)
	}

	g3 := g.Group("/api/v1/goods")
	{
		g3.POST("/name", goods.GetGoodsInStockByName)
		g3.POST("/num", goods.GetGoodsInStockByNum)
		g3.POST("/supplier", goods.GetGoodsInStockBySupplier)
	}

	g4 := g.Group("/api/v1/orders")
	{
		g4.POST("/", order.CreateOrder)
		g4.POST("/time", order.GetOrdersByTime)
		g4.POST("/user", order.GetOrdersByUser)
		g4.POST("/return", order.CreateReturnOrder)
	}
	g5 := g.Group("/api/v1/profit")
	{
		g5.POST("/", profit.GetProfitByTime)
		g5.POST("/name", profit.GetProfitByTimeAndName)

	}

}

// 配置 404 路由
func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

}
