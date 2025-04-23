package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r Router) AddReviewRouter(rg *gin.RouterGroup) {

	reviewCont := r.deps.ReviewController
	reviewRouter := rg.Group("review")

	reviewRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), reviewCont.CreateReview)
	reviewRouter.GET("/", reviewCont.GetReviews)
	reviewRouter.GET("/:review_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), reviewCont.GetReviewByID)
	reviewRouter.GET("/seller/:seller_id", reviewCont.GetReviewsBySellerID)
	reviewRouter.GET("/buyer/:buyer_id", reviewCont.GetReviewsByBuyerID)
	reviewRouter.PUT("/:review_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), reviewCont.UpdateReview)
	reviewRouter.DELETE("/:review_id", reviewCont.DeleteReview)

}
