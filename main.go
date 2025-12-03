package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Struktur data sayuran
type Vegetable struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Buat router dengan logger dan recovery middleware
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// Custom log format
		return fmt.Sprintf("[%s] %s - %s %s %d %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))
	r.Use(gin.Recovery())

	// Data sayuran
	vegetables := []Vegetable{
		{ID: 1, Name: "Wortel"},
		{ID: 2, Name: "Kubis"},
		{ID: 3, Name: "Bayam"},
		{ID: 4, Name: "Tomat"},
		{ID: 5, Name: "Timun"},
	}

	// Endpoint GET /vegetables
	r.GET("/vegetables", func(c *gin.Context) {
		log.Println("Endpoint /vegetables dipanggil")
		c.JSON(http.StatusOK, vegetables)
	})

	port := ":8081"
	log.Println("Menjalankan server di port", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
