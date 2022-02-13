package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xcbeyond/pkg/gin/common"
)

func ErrorMiddleware() gin.HandlerFunc {
	return errorHandler(gin.ErrorTypeAny)
}

func errorHandler(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		var err *common.ResponseRes
		detectedErrs := c.Errors.ByType(errType)
		if len(detectedErrs) > 0 {
			for _, v := range detectedErrs {
				e := v.Err
				switch e.(type) {
				case *common.ResponseRes:
					err = e.(*common.ResponseRes)
				default:
					err = &common.ResponseRes{Code: http.StatusText(http.StatusInternalServerError), Message: e.Error()}
				}
			}
			c.IndentedJSON(http.StatusInternalServerError, err)
			c.Abort()
		}
	}
}
