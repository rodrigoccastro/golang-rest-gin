package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func getResultDataItem(data gin.H) gin.H {
	return gin.H{
			"time": time.Now(),
			"data": data,
		}
}

func getResultDataList(data []gin.H) gin.H {
	return gin.H{
			"time": time.Now(),
			"data": data,
		}
}

func getResultError(context string, err error) gin.H {
	fmt.Println("Error in "+context)
	fmt.Println(err)
	
	return gin.H{
			"time": time.Now(),
			"error": err,
		}
}