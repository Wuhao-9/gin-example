package middleware

import (
	"github/Wuhao-9/go-gin-example/pkg/ec"
	"github/Wuhao-9/go-gin-example/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := ec.SUCCESS
		token := ctx.Query("token")
		if token == "" {
			code = ec.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = ec.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = ec.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != ec.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  ec.GetMsg(code),
				"data": nil,
			})
			ctx.Abort() // 终止当前路由
		}
		ctx.Next()
	}
}
