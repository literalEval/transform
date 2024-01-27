package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func transform(ctx *gin.Context) {
	cmd := exec.Command("primitive", strings.Fields("-i res/input.png -o res/output.png -n 250 -m 0")...)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Err: ", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Done !!",
	})
}

func main() {
	router := gin.Default()

	router.GET("/transform", func(ctx *gin.Context) {
		transform(ctx)
	})

	router.Run()
}
