package handler

import (
	"code.aliyun.com/zmdev/wechat_rank/handler/middleware"
	"code.aliyun.com/zmdev/wechat_rank/server"
	"github.com/gin-gonic/gin"
	gerrorsGin "github.com/zm-dev/gerrors/gin"
	"net/http"
	"strconv"
)

const (
	ServiceName = "wechat_rank"
)

func CreateHTTPHandler(svr *server.Server) http.Handler {

	if svr.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	wechatHandler := NewWechat()
	categoryHandler := NewCategory()
	rankHandler := NewRank()
	authHandler := NewAuth()
	articleHandler := NewArticle()
	exportHandler := NewExport()

	router := gin.Default()
	router.Use(middleware.ServiceMiddleware(svr.Service))
	router.Use(gerrorsGin.NewHandleErrorMiddleware(svr.ServiceName))
	authRouter := router.Group("/auth")
	// 登录
	authRouter.POST("/login", authHandler.Login)
	// 注册
	authRouter.POST("/register", authHandler.Register)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// 退出登录
		authorized.GET("/logout", authHandler.Logout)
		// 添加公众号
		authorized.POST("/wechat", wechatHandler.Create)
		// 删除公众号
		authorized.DELETE("/wechat/:id", wechatHandler.Delete)

		// 创建分类
		authorized.POST("/category", categoryHandler.Create)
		// 删除分类
		authorized.DELETE("/category/:id", categoryHandler.Delete)
		// 更新分类
		authorized.PUT("/category/:id", categoryHandler.Update)
	}
	// 所有公众号列表(好像没有用)
	//router.GET("/wechat", wechatHandler.List)
	// 指定分类下的公众号列表
	router.GET("/category/:id/wechat", wechatHandler.ListByCategory)
	// 分类列表
	router.GET("/category", categoryHandler.List)
	// 分类详情
	router.GET("/category/:id", categoryHandler.Show)

	// 排名
	// 获取日期
	router.GET("/rank/date", rankHandler.RankList)
	// 公众号排名
	router.GET("/rank/account", rankHandler.AccountRank)
	// 文章排名(日期区间随意)
	router.GET("/rank/article", rankHandler.ArticleRank)

	// 指定公众号的所有文章
	router.GET("/article", articleHandler.List)

	// 导出公众号排名
	router.GET("/export/account", exportHandler.AccountRank)
	// 导出文章排名
	router.GET("/export/article", exportHandler.ArticleRank)
	return router
}

func getLimitAndOffset(c *gin.Context) (limit, offset int) {
	var err error
	limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	if limit > 200 {
		limit = 200
	}
	offset, err = strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	return limit, offset
}
