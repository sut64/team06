package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sut64/team06/backend/controller"
	"github.com/sut64/team06/backend/entity"
	"github.com/sut64/team06/backend/middlewares"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// User Routes ** optional

			// ManageSalary Routes

			// ManageWorkSchedule Routes
			protected.POST("/manage/create", controller.CreateMangeWorkTime)
			protected.GET("/manage/day/all", controller.GetAllDay)
			protected.GET("/manage/week/all", controller.GetAllWeekly)
			protected.GET("/manage/month/all", controller.GetAllMonth)
			protected.GET("/manage/work/all", controller.GetAllWorkTime)
			protected.GET("/manage/employee/all", controller.GetAllEmployee)
			protected.GET("/manage/detail/all", controller.GetAllDetail)
			protected.GET("/manage/pos/all", controller.GetAllPosition)
			protected.GET("/manage/all", controller.GetAllManageWorkTime)

			// ProductStock Routes
			protected.GET("/productstock", controller.ListProductstock)
			protected.GET("/typeproduct", controller.ListTypeproduct)
			protected.GET("/product", controller.ListProduct)
			protected.GET("/shelfstore", controller.ListShelfstore)
			protected.POST("/productstock", controller.CreateProductstock)
			protected.POST("/product", controller.CreateProduct)

			// ManagePromotion Routes
			protected.GET("/listpromotion", controller.ListPromotion)
			protected.GET("/getnamepromotion", controller.GetNamePromotion)
			protected.GET("/getpromotionperiod", controller.GetPromotionperiod)
			protected.GET("/getpromotiontype", controller.GetPromotionType)
			protected.POST("/createpromotion", controller.CreateManagePromotion)

			// PremiumMember Routes

			// PurchaseOrder Routes
			protected.GET("/payment-methods", controller.ListPaymentMethod)
			protected.GET("/order-history/:id", controller.ListPurchaseOrder)
			protected.POST("/purchase-order", controller.CreatePurchaseOrder)
		}
	}

	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}
