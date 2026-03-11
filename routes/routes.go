package routes

import (
	"payment-service/controllers"
	"payment-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/payment")
	{
		// 公开端点
		api.POST("/apple-pay/session", controllers.CreatePaymentSession)
		api.GET("/status/:orderId", controllers.GetPaymentStatus)

		// 需要认证的端点
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/apple-pay/process", controllers.ProcessApplePayment)
			protected.POST("/refund/:orderId", controllers.RefundPayment)
		}
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "payment-service"})
	})
}
