package routes

import (
	"github.com/stivan622/kiokumushi-api/app/http/controllers/web"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func MapWebRoutes(router *gin.Engine) {
	// マッピング
	mapping(router)

	// HTMLレンダラーの設定
	settingHTMLRenderer(router)
}

func mapping(router *gin.Engine) {
	// トップページ
	homeController := &web.HomeController{}
	{
		router.GET("/", homeController.Index)
	}

	// ヘルスチェック
	healthController := &web.HealthController{}
	{
		router.GET("/health", healthController.Index)
	}

	// ムシ
	mushiController := &web.MushiController{}
	{
		router.GET("/mushi", mushiController.Index)
	}
}

func settingHTMLRenderer(router *gin.Engine) {

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "../resources/views/web/",
		Extension: ".tpl",
		Master:    "layouts/main",
	})

}
