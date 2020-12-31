package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context)  {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"sitename": "Superset",
		"siteurl": "http://superset.xxxx.com.cn",
	})
}