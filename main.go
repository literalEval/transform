package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func transform(ctx *gin.Context) {
	file, err := ctx.FormFile("input_file")

	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Err !!",
		})

		return
	}

	n := ctx.PostForm("cnt")
	m := ctx.PostForm("types")

	log.Println(file.Filename, n, m)

	ctx.SaveUploadedFile(file, "res/"+file.Filename)

	cmd := exec.Command("primitive", strings.Fields("-i res/"+file.Filename+" -o res/output.jpg -n "+n+" -m "+m)...)
	err = cmd.Run()

	if err != nil {
		log.Println("Err: ", err)
	}

	// ctx.File("res/output.jpg")

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"output": "res/output.jpg",
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/home.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	router.GET("/res/output.jpg", func(ctx *gin.Context) {
		ctx.File("res/output.jpg")
	})

	router.POST("/transform", func(ctx *gin.Context) {
		transform(ctx)
	})

	router.Run()
}
