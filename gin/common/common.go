package common

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	ginI18n "github.com/xcbeyond/pkg/gin/i18n"
)

type ResponseRes struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (t *ResponseRes) Error() string {
	return t.Message
}

func Response(data interface{}) *ResponseRes {
	return &ResponseRes{
		Code:    "0000",
		Message: "Success",
		Data:    data,
	}
}

func Error(messageID string, param interface{}) *ResponseRes {
	message := ginI18n.MustGetMessage(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: param,
	})
	return &ResponseRes{
		Code:    messageID,
		Message: message,
	}
}

// func Success(c *gin.Context, data interface{}) {
// 	message := ginI18n.MustGetMessage(&i18n.LocalizeConfig{
// 		MessageID: "ok",
// 	})

// 	obj := &ResponseResult{
// 		Code:    http.StatusOK,
// 		Message: message,
// 		Data:    data,
// 	}

// 	c.JSON(http.StatusOK, obj)
// }

// func Failed(c *gin.Context, messageID string, param interface{}) {
// 	message := ginI18n.MustGetMessage(&i18n.LocalizeConfig{
// 		MessageID:    messageID,
// 		TemplateData: param,
// 	})

// 	obj := &ResponseResult{
// 		Code:    http.StatusInternalServerError,
// 		Message: message,
// 	}

// 	c.IndentedJSON(http.StatusInternalServerError, obj)
// 	c.Abort()
// }
