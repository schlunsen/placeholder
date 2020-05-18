package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"strconv"

	"github.com/anthonynsimon/bild/transform"
	"github.com/julienschmidt/httprouter"
)

func ImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var imgjpg, _ = LoadImage("img/test.jpeg")

	width, err := strconv.Atoi(ps.ByName("width"))
	height, err := strconv.Atoi(ps.ByName("height"))

	if err != nil {
		fmt.Println("HEST")
	}
	ip := GetIP(r)

	msg := fmt.Sprintf("Width: %v, Height: %v Ip: ‰v", width, height, ip)

	log.Println(msg)

	result := transform.Resize(imgjpg, width, height, transform.Linear)

	var img image.Image = result

	writeImage(w, &img)

}
