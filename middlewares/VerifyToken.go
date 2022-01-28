package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")
		c.Next()
		latency := time.Since(t)
		log.Println(latency)
		log.Printf("ClientIP: %s\n", c.ClientIP())
	}
}
