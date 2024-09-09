package main

import (
	"log"
	"net/http"
)

func main() {
	// Servir la carpeta 'static' para archivos como CSS
	fs := http.FileServer(http.Dir("./main/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // Sirve archivos estáticos desde '/static/'

	// Ruta para servir el archivo index.html en la raíz
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./main/static/index.html")
	})

	// Configuración del servidor
	port := "8080"
	log.Printf("Server listening on http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
