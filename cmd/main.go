package main

import (
	"github.com/gin-gonic/gin"
	"short.cat/internal/router"
)

func main() {
	r := gin.Default()
	r = router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
