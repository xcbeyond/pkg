package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/xcbeyond/pkg/i18n"
	"golang.org/x/text/language"
)

func main() {
	// 初始化gin对象
	route := gin.Default()

	route.Use(i18n.Localize(i18n.WithBundle(&i18n.BundleCfg{
		RootPath:         "./error",
		AcceptLanguage:   []language.Tag{language.Chinese, language.English},
		DefaultLanguage:  language.English,
		UnmarshalFunc:    json.Unmarshal,
		FormatBundleFile: "json",
	})))

	// 设置一个get请求，其URL为/hello，并实现简单的响应
	route.GET("/hello", func(c *gin.Context) {
		c.JSON(200, i18n.MustGetMessage("hello"))
	})

	// 启动服务
	route.Run()
}
