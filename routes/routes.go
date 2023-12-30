package routes

import (
	"golang-web/models"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func HomeRoutes(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("public", "index.html"), path.Join("public", "layout.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "Siksa Neraka",
	// 	"content": "Berhentilah berbuat dosa sesungguh nya siksa neraka itu amat sangat kejam",
	// }

	data := []models.Product{
		{ID: 1, Name: "Vespa Sprint", Price: 4800000000, Stock: 1},
		{ID: 1, Name: "Vespa Primavera", Price: 4400000000, Stock: 9},
		{ID: 1, Name: "Vespa Super", Price: 4200000000, Stock: 13},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AboutRoutes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini Halaman About"))
}

func ContactRoutes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ini Halaman Contact"))
}

func ProductRoutes(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("public", "product.html"), path.Join("public", "layout.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content": idNumb,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
