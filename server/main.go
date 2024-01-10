package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFS("/", http.Dir("web"))

	log.Printf("Listening on port 8000")
	log.Fatal(r.Run(":8000"))
}
