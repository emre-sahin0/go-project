package handlers

import (
	"fmt"
	"net/http"
	"strings"
)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}




func DynamicHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/merhaba/")
	name := strings.Title(path)
	if name == "" {
		name = "Ziyaretçi"
	}
	fmt.Fprintf(w, "Merhaba, %s!", name)
}
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Formu göster
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="tr">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Form Sayfası</title>
			</head>
			<body>
				<h1>Adınızı Girin</h1>
				<form action="/form" method="post">
					<label for="name">Adınız:</label>
					<input type="text" id="name" name="name">
					<button type="submit">Gönder</button>
				</form>
			</body>
			</html>
		`)
	} else if r.Method == http.MethodPost {
		// Formdan gelen veriyi işle
		name := r.FormValue("name")
		if name == "" {
			name = "Bilinmeyen"
		}
		fmt.Fprintf(w, "Merhaba, %s! Formunuzu aldık.", name)
	}
}
