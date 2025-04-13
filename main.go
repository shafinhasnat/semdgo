package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var mds = `# header

Sample text.

[link](http://example.com)
`

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
	// md := []byte(mds)
	// html := mdToHTML(md)

	// fmt.Printf("--- Markdown:\n%s\n\n--- HTML:\n%s\n", md, html)
	http.HandleFunc("/", handle)
	http.ListenAndServe(":80", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	path := clickedhyperlink(r)
	log.Printf("path: %s\n", path)
	content, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Unable to read markdown file", http.StatusInternalServerError)
		return
	}
	html := mdToHTML(content)
	// Log clicked hyperlinks
	fmt.Fprintf(w, "%s", html)
}

func clickedhyperlink(r *http.Request) string {
	basepath := "./content"
	link := r.URL.Path
	if !strings.Contains(link, "favicon.ico") {
		if link == "/" {
			path := basepath + "/README.md"
			return path
		}
		path := basepath + link
		return path
	}
	return basepath + "/README.md"
}
