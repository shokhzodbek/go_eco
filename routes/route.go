package routes

import "github.com/gin-gonic/gin"


func UserRoutes(route *gin.Engine)  {
	route.POST("/users/signup")
	route.POST("/users/login")
	route.POST("/admin/addproduct")
	route.GET("/users/productview")
	route.GET("/users/serach")
	
}


