package server

import (
	"test_restapi/controller"
	"test_restapi/db"

	"github.com/gin-gonic/gin"
)

// Handler return nil
func Handler(port string) {
	db.Open()
	defer db.Close()
	r := gin.Default()

	r.GET("/medicine", controller.GetMedicine)
	r.POST("/medicine", controller.CreateMedicine)
	r.PUT("/medicine/:id", controller.UpdateMedicine)
	r.DELETE("/medicine/:id", controller.DeleteMedicine)

	r.Run(":" + port)
}
