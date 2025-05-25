package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct Mahasiswa
type Mahasiswa struct {
	NPM   int    `json:"npm"`
	Nama  string `json:"nama"`
	Kelas string `json:"kelas"`
}

func main() {
	router := gin.Default()

	// Endpoint root
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Selamat datang di API Mahasiswa menggunakan Gin!",
			"author":  "Savana Zahra Humaira",
		})
	})

	// Endpoint GET /mahasiswa
	router.GET("/mahasiswa", func(c *gin.Context) {
		mahasiswa := Mahasiswa{
			NPM:   714230042,
			Nama:  "Savana Zahra Humaira",
			Kelas: "TI-3A",
		}
		c.JSON(http.StatusOK, mahasiswa)
	})

	// Endpoint POST /mahasiswa
	router.POST("/mahasiswa", func(c *gin.Context) {
		var input Mahasiswa

		// Parsing JSON dari body request
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Data tidak valid",
			})
			return
		}

		// Balikkan data yang diterima
		c.JSON(http.StatusOK, gin.H{
			"message":   "Data mahasiswa berhasil diterima",
			"mahasiswa": input,
		})
	})

	// Jalankan server
	router.Run(":8080")
}

.