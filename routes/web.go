package routes

import (
	"kiokumushi-api/app/http/controllers/web"

	"github.com/gin-gonic/gin"
)

func MapWebRoutes(router *gin.Engine) {
	// マッピング
	mapping(router)

	// ビューを読み込み
	router.LoadHTMLGlob("../resources/views/web/**/*")
}

func mapping(router *gin.Engine) {
	// トップページ
	homeController := new(web.HomeController)
	{
		router.GET("/", homeController.Index)
	}

	// ヘルスチェック
	healthController := new(web.HealthController)
	{
		router.GET("/health", healthController.Index)
	}

	// ムシ
	mushiController := new(web.MushiController)
	{
		router.GET("/mushi", mushiController.Index)
	}
}
