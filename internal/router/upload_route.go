package router

import (
	"github.com/gin-gonic/gin"
)

func (r Router) AddUploadRoute(rg *gin.RouterGroup) {

	uploadCont := r.deps.S3Controller
	uploadRouter := rg.Group("upload")

	uploadRouter.POST("/", uploadCont.UploadFile)

}
