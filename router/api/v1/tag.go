package v1

import (
	"github/Wuhao-9/go-gin-example/models/dao"
	"github/Wuhao-9/go-gin-example/pkg/ec"
	"github/Wuhao-9/go-gin-example/pkg/setting"
	"github/Wuhao-9/go-gin-example/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取多个文章标签
func GetTags(ctx *gin.Context) {
	constraints := make(map[string]interface{})
	data := make(map[string]interface{})

	if name := ctx.Query("name"); name != "" {
		constraints["name"] = name
	}

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		constraints["state"] = state
	}

	code := ec.SUCCESS

	data["lists"] = dao.GetTags(util.GetOffset(ctx), setting.PageSize, constraints)
	data["total"] = dao.GetTagTotal(constraints)

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ec.GetMsg(code),
		"data": data,
	})
}

// 新增文章标签
func AddTag(ctx *gin.Context) {
	name := ctx.Query("name")
	state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()
	createdBy := ctx.Query("created_by")

	valid := &validation.Validation{} // 创建一个Validation对象用来验证
	var res *validation.Result        // 创建一个验证结果对象用来保存结果
	res = valid.Required(name, "required <name>").Message("名称不能为空")
	if res.Ok {
		valid.MaxSize(name, 20, "required size <name>").Message("名称最长为20个字符")
	}
	res = valid.Required(createdBy, "required <created_by>").Message("创建人不能为空")
	if res.Ok {
		valid.MaxSize(createdBy, 20, "required size <created_by>").Message("创建人最长为20字符")
	}
	res = valid.Range(state, 0, 1, "range <state>").Message("状态可选值为 0 or 1")

	code := ec.INVALID_PARAMS
	if !valid.HasErrors() {
		if !dao.IsExist_tag_by_name(name) {
			code = ec.SUCCESS
			dao.AddTag(name, state, createdBy)
		} else {
			code = ec.ERROR_EXIST_TAG
		}
	} else {
		for _, v := range valid.Errors {
			log.Println(v.Key, ":", v.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": ec.GetMsg(code),
		"data":    nil,
	})
}

// 修改文章标签
func EditTag(ctx *gin.Context) {
	var id int = -1
	if v, err := com.StrTo(ctx.Param("id")).Int(); err == nil {
		id = v
	}

	valid := validation.Validation{}
	var res *validation.Result
	var state int = -1
	if v := ctx.Query("state"); v != "" {
		if tmp, err := com.StrTo(v).Int(); err == nil {
			state = tmp
		}
		valid.Range(state, 0, 1, "range <state>").Message("状态可选值为 0 or 1")
	}

	name := ctx.Query("name")
	res = valid.Required(name, "required <name>").Message("名称不能为空")
	if res.Ok {
		valid.MaxSize(name, 20, "required size <name>").Message("名称最长为20个字符")
	}

	modified_by := ctx.Query("modified_by")
	res = valid.Required(modified_by, "required <name>").Message("修改人不能为空")
	if res.Ok {
		valid.MaxSize(modified_by, 20, "required size <modified_by>").Message("名称最长为20个字符")
	}

	code := ec.INVALID_PARAMS
	if !valid.HasErrors() {
		if dao.IsExist_tag_by_id(id) {
			code = ec.SUCCESS
			data := make(map[string]interface{})
			data["modified_by"] = modified_by
			data["name"] = name
			data["state"] = state
			dao.EditTag(id, data)
		} else {
			code = ec.ERROR_NOT_EXIST_TAG
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ec.GetMsg(code),
		"data": nil,
	})
}

// 删除文章标签
func DeleteTag(ctx *gin.Context) {
	var target_id int = -1
	if v, err := com.StrTo(ctx.Param("id")).Int(); err == nil {
		target_id = v
	}

	valid := &validation.Validation{}
	valid.Min(target_id, 1, "Min <target_id>").Message("ID必须大于0")

	code := ec.INVALID_PARAMS
	if !valid.HasErrors() {
		if dao.IsExist_tag_by_id(target_id) {
			code = ec.SUCCESS
			dao.DeleteTag(target_id)
		} else {
			code = ec.ERROR_NOT_EXIST_TAG
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ec.GetMsg(code),
		"data": nil,
	})
}
