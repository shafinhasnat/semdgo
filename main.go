package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":80", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	path := clickedhyperlink(r)
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
	html := mdToHTML(content)
	tmpl, err := template.ParseFiles("./templates/200")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, map[string]interface{}{
		"Content": string(html),
	})

}

func clickedhyperlink(r *http.Request) string {
	basepath := "/var/semdgo/content"
	link := r.URL.Path
	if link == "/" {
		path := basepath + "/README.md"
		return path
	}
	path := basepath + link
	return path
}
