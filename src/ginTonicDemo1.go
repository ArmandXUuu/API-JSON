package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RepArmandXU() {
	fmt.Println("一个网站上的教程，简单")

	r := gin.Default()
	r.GET("/ArmandXU", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello ArmandXU !",
		})
	})

	r.Run()
}
