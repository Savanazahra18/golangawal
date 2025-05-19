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
			<style>
				body {
					font-family: Arial, sans-serif;
					margin: 0;
					padding: 0;
				}
				.navbar {
					background-color: #333;
					overflow: hidden;
				}
				.navbar a {
					float: left;
					display: block;
					color: #f2f2f2;
					text-align: center;
					padding: 14px 16px;
					text-decoration: none;
				}
				.navbar a:hover {
					background-color: #ddd;
					color: black;
				}
				.container {
					padding: 20px;
				}
			</style>
		</head>
		<body>
			<div class="navbar">
				<a href="/">Beranda</a>
				<a href="/tentang">Tentang</a>
				<a href="/kontak">Kontak</a>
			</div>
			<div class="container">
				<h1>Selamat Datang di Website Saya</h1>
				<p>Ini adalah halaman index yang dibuat dengan Go.</p>
			</div>
		</body>
		</html>
	`)
}
.

