package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start Listning...")
	r := gin.Default()
	r.Run()
}
