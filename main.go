// Package main adalah entry point atau titik awal dari aplikasi.
package main

// Mengimpor library standar (fmt) dan modul internal (config, database, route) 
// untuk membangun pondasi aplikasi sebelum dijalankan.
import (
	// librari standar
	"fmt"

	// modul internal
	"latihan/config"
	"latihan/database"
	"latihan/route"
)

func main() {
	// 1. Load data dari .env
	cfg := config.LoadConfig()

	// 2. Koneksi ke database
	database.ConnectMySQL(cfg)
	// database.ConnectPostgres(cfg)
    
	// 3. Jalankan server
	r := routes.SetupRouter()
	r.Run(":7000")

	// 4. Tampilkan pesan
	fmt.Println("Server berjalan di http://localhost:7000")
}