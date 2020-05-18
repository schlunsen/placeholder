package main

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	log.Println("Image loaded")
	switch filepath.Ext(path) {
	case ".png":
		return png.Decode(file)
	case ".jpeg":
		return jpeg.Decode(file)
	case ".gif":
		return gif.Decode(file)
	default:
		return nil, errors.New("Unkown file format.")
	}
	return jpeg.Decode(file)
}

func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
