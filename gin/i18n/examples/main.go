package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xcbeyond/pkg/gin/common"
	"github.com/xcbeyond/pkg/gin/error"
	ginI18n "github.com/xcbeyond/pkg/gin/i18n"
	"golang.org/x/text/language"
)

func main() {
	// 初始化gin对象
	route := gin.Default()

	route.Use(error.ErrorMiddleware())
	route.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		RootPath:         "./error",
		AcceptLanguage:   []language.Tag{language.Chinese, language.English},
		DefaultLanguage:  language.English,
		UnmarshalFunc:    json.Unmarshal,
		FormatBundleFile: "json",
	})))

	// 设置一个get请求，其URL为/hello，并实现简单的响应
	// route.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, ginI18n.Message("hello", nil))
	// })

	route.GET("/hello", func(c *gin.Context) {
		// c.JSON(http.StatusOK, ginI18n.MustGetMessage(&i18n.LocalizeConfig{
		// 	MessageID: "helloWithName",
		// 	TemplateData: map[string]string{
		// 		"name": c.Param("name"),
		// 	},
		// }))
		if c.Query("name") == "" {
			// c.Error(fmt.Errorf("1111"))
			c.Error(common.Error("1001", nil))
			return
		}

		resData := c.Query("name")
		c.JSON(http.StatusOK, common.Response(resData))
	})

	// 启动服务
	route.Run()
}
