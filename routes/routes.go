package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"bluebell/settings"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 每两秒填充1个令牌，请求拿到桶中令牌才能获取响应，如果拿不到就获取不到响应
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(time.Second*2, 1))
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignupHandler)
	v1.POST("/login", controller.LoginHandler)
	// JWT中间对use下面的中间件才有作用
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/index", func(c *gin.Context) {
			c.String(http.StatusOK, settings.Conf.Version)
		})
		// 获取社区分类
		v1.GET("/community", controller.CommunityListHandler)
		// 通过id查询社区分类详情
		v1.GET("/community/:id", controller.CommunityDetailByIDHandler)
		// 用户新增帖子功能接口
		v1.POST("/post", controller.CreatePostHandler)
		// 通过帖子id查询帖子的详情，包括用户该帖子的作者信息以及社区帖子分类信息
		v1.GET("/post/:id", controller.PostDetailByIDHandler)
		// 实现分页获取帖子
		v1.GET("/posts", controller.PostListHandler)
		// 实现根据前端传来的参数，按时间排序返回或者按分数排序返回+返回帖子的投票数
		v1.GET("/posts_order", controller.PostListOrderHandler)
		// 给帖子投票
		v1.POST("/vote", controller.PostVoteHandler)
	}

	return r
}
