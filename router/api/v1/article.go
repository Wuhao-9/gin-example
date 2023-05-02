package v1

import (
	"github/Wuhao-9/go-gin-example/models"
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

func GetArticles(ctx *gin.Context) {
	data := make(map[string]interface{})
	constraint := make(map[string]interface{})

	valid := validation.Validation{}
	if state := ctx.Query("state"); state != "" {
		state, err := com.StrTo(state).Int()
		if err != nil {
			if valid.Range(state, 0, 1, "Range <state>").Message("状态只允许为0和1").Ok {
				constraint["state"] = state
			}

		}
	}
	if tag_id := ctx.Query("tag_id"); tag_id != "" {
		tag_id, err := com.StrTo(tag_id).Int()
		if err != nil {
			if valid.Min(tag_id, 0, "目标标签ID必须大于0").Ok {
				constraint["tag_id"] = tag_id
			}
		}
	}

	code := ec.INVALID_PARAMS
	if !valid.HasErrors() {
		code = ec.SUCCESS
		data["list"], _ = dao.GetArticles(util.GetOffset(ctx), setting.PageSize, constraint)
		data["total"] = dao.GetArticleTotal(constraint)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ec.GetMsg(code),
		"data": data,
	})
}

func GetArticle(ctx *gin.Context) {
	tar_id, err := com.StrTo(ctx.Param("id")).Int()
	if err != nil {
		tar_id = -1
	}

	valid := validation.Validation{}
	valid.Min(tar_id, 0, "Min <target_id>").Message("文章ID必须大于0")

	code := ec.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if dao.IsExist_article_by_id(tar_id) {
			data = dao.GetArticle(tar_id)
			code = ec.SUCCESS
		} else {
			code = ec.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("[validation] err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ec.GetMsg(code),
		"data": data,
	})
}

func AddArticle(ctx *gin.Context) {
	title := ctx.Query("title")
	content := ctx.Query("content")
	desc := ctx.Query("desc")
	created_by := ctx.Query("created_by")
	state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()
	tag_id, err := com.StrTo(ctx.Query("tag_id")).Int()
	if err != nil {
		tag_id = -1
	}

	valid := validation.Validation{}
	valid.Min(tag_id, 0, "Min <tag_id>").Message("标签ID必须大于0")
	if valid.Required(title, "Required <title>").Message("文章标题不能为空").Ok {
		valid.MaxSize(title, 50, "MaxSize <title>").Message("文章标题应在50字符之内")
	}
	valid.Required(content, "Required <content>").Message("文章内容不能为空")
	if valid.Required(created_by, "Required <created_by>").Message("创建人不能为空").Ok {
		valid.MaxSize(content, 20, "MaxSize <create_by>").Message("创建人姓名应在20字符之内")
	}
	valid.Range(state, 0, 1, "Range <state>").Message("状态只允许0或1")

	code := ec.INVALID_PARAMS
	if !valid.HasErrors() {
		if dao.IsExist_tag_by_id(tag_id) {
			record_info := models.Article{
				Title:     title,
				State:     state,
				CreatedBy: created_by,
				TagID:     tag_id,
				Desc:      desc,
				Content:   content,
			}
			if dao.AddArticle(&record_info) {
				code = ec.SUCCESS
			} else {
				code = ec.ERROR
			}
		} else {
			code = ec.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("[validation] err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ec.GetMsg(code),
		"data": nil,
	})
}

func DeleteArticle(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 0, "id").Message("目标ID必须大于0")

	code := ec.INVALID_PARAMS
	if !valid.HasErrors() {
		if dao.IsExist_article_by_id(id) {
			dao.DeleteArticle(id)
			code = ec.SUCCESS
		} else {
			code = ec.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  ec.GetMsg(code),
		"data": nil,
	})
}
