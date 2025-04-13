package handler

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/shafinhasnat/semdgo/internal/renderer"
	"github.com/shafinhasnat/semdgo/internal/utils"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	path := utils.ClickedHyperlink(r)
	log.Printf("Status: %d | path: %s\n", http.StatusOK, path)
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Status: %d | path: %s | Error: %s\n", http.StatusNotFound, path, err)
		http.Error(w, "Unable to read markdown file", http.StatusNotFound)
		return
	}
	if !strings.HasSuffix(path, ".md") {
		http.ServeFile(w, r, path)
		return
	}
	html := renderer.MDtoHTML(content)
	tmpl, err := template.ParseFiles("./templates/200")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, map[string]interface{}{
		"Content": string(html),
		"Path":    strings.TrimPrefix(path, "/var/semdgo/content"),
	})
}
