package main

import (
	"golang-web/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", routes.HomeRoutes)
	mux.HandleFunc("/about", routes.AboutRoutes)
	mux.HandleFunc("/contact", routes.ContactRoutes)
	mux.HandleFunc("/product", routes.ProductRoutes)

	fileServer := http.FileServer(http.Dir("css"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("server run on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
