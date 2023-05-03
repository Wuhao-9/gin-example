package v1

import (
	"github/Wuhao-9/go-gin-example/models/dao"
	"github/Wuhao-9/go-gin-example/pkg/ec"
	"github/Wuhao-9/go-gin-example/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Account string `valid:"Required;MaxSize(20)"`
	Passwd  string `valid:"Required;MaxSize(20)"`
}

func GetAuth(ctx *gin.Context) {
	account := ctx.Query("account")
	passwd := ctx.Query("passwd")

	valid := validation.Validation{}
	cur_user := auth{Account: account, Passwd: passwd}

	ok, _ := valid.Valid(cur_user)
	data := make(map[string]interface{})
	code := ec.INVALID_PARAMS
	if ok {
		isExist := dao.CheckAuth(account, passwd)
		if isExist { // 当前账户存在
			token, err := util.GenerateToken(account, passwd)
			if err != nil {
				code = ec.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = ec.SUCCESS
			}
		} else { // 当前账户不存在
			code = ec.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ec.GetMsg(code),
		"data": data,
	})
}
