package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)


func main() {
	port:=os.Getenv("PORT")
	if port=="" {
		port="8000"
	}
    router := gin.New()
	router.Use(gin.Logger())
	router.GET("/addcart")
	router.GET("/removeitem")
	router.GET("/cartcheckout")
	router.GET("/instantbuy")

	log.Fatal(router.Run(":"+ port))

}