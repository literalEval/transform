package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func getM(strm string) string {
	str_m := map[string]string{
		"Combo":           "0",
		"Triangle":        "1",
		"Rectangle":       "2",
		"Ellipse":         "3",
		"Circle":          "4",
		"Rotate Direct":   "5",
		"Beziers":         "6",
		"Rotated Eclipse": "7",
		"Polygon":         "8",
		"":                "0",
	}

	return str_m[strm]
}

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
	m := getM(ctx.PostForm("types"))

	log.Println(file.Filename, n, m)

	ctx.SaveUploadedFile(file, "res/input.jpg")

	cmd := exec.Command("primitive", strings.Fields("-i res/"+file.Filename+" -o res/output.jpg -n "+n+" -m "+m)...)
	err = cmd.Run()

	if err != nil {
		log.Println("Err: ", err)
	}


	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"input":  "input.jpg",
		"output": "output.jpg",
	})
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/home.html")

	router.StaticFile("input.jpg", "res/input.jpg")
	router.StaticFile("output.jpg", "res/output.jpg")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	router.POST("/transform", func(ctx *gin.Context) {
		transform(ctx)
	})

	router.Run()
}
