package util

import (
	"github/Wuhao-9/go-gin-example/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetOffset(ctx *gin.Context) (result int) {
	page := com.StrTo(ctx.Query("page")).MustInt() // TODO: 此处进行错误log处理
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return
}
