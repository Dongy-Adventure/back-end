package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r Router) AddUploadRoute(rg *gin.RouterGroup) {

	uploadCont := r.deps.S3Controller
	uploadRouter := rg.Group("upload")

	uploadRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), uploadCont.UploadFile)

}
