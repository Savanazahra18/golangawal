package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Rute untuk halaman index
	http.HandleFunc("/", indexHandler)

	// Menentukan port server
	port := ":8080"
	fmt.Printf("Server berjalan di http://localhost%s\n", port)

	// Menjalankan server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}

// Fungsi handler untuk halaman index
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html lang="id">
		<head>
			<meta charset="UTF-8">
			<title>Halaman Index</title>
		</head>
		<body>
			<h1>Selamat Datang di Website Saya</h1>
			<p>Ini adalah halaman index yang dibuat dengan Go.</p>
		</body>
		</html>
	`)
}
