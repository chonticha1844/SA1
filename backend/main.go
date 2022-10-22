package main

import (
	"github.com/chonticha1844/sa-65-example/controller"
	"github.com/chonticha1844/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	router := gin.Default()
	router.Use(CORSMiddleware())

	// Employee Routes
	router.GET("/employees", controller.ListEmployees)
	//router.GET("/employee/:id", controller.ListEmployees)
	router.GET("/employee/:id", controller.GetEmployee)
	router.POST("/employees", controller.CreateEmployee)
	router.PATCH("/employees", controller.UpdateEmployee)
	router.DELETE("/employees/:id", controller.DeleteEmployee)

	// Typeproduct Routes
	router.GET("/genders", controller.ListGenders)
	router.GET("/gender/:id", controller.GetGender)
	router.POST("/genders", controller.CreateGender)
	router.PATCH("/genders", controller.UpdateGender)
	router.DELETE("/genders/:id", controller.DeleteGender)

	// Manufacturer Routes
	router.GET("/provinces", controller.ListProvinces)
	router.GET("/province/:id", controller.GetProvince)
	router.POST("/provinces", controller.CreateProvince)
	router.PATCH("/provinces", controller.UpdateProvince)
	router.DELETE("/provinces/:id", controller.DeleteProvince)

	// Product Routes
	router.GET("/members", controller.ListMembers)
	router.GET("/member/:id", controller.GetMember)
	router.POST("/members", controller.CreateMember)
	router.PATCH("/members", controller.UpdateMember)
	router.DELETE("/members/:id", controller.DeleteMember)

	// Run the server

	router.Run()

}
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
