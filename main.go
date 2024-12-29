package main

import (
	"fmt"
	"net/http"

	"my-go-project/handlers" // Modül adını kullanarak doğru import
)

func main() {
	// Statik dosya sunumu
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Ana sayfa
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/form", handlers.FormHandler)




	// Dinamik endpoint
	http.HandleFunc("/merhaba/", handlers.DynamicHandler)

	// Sunucuyu başlat
	fmt.Println("Sunucu çalışıyor: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Sunucu başlatılamadı:", err)
	}
}
