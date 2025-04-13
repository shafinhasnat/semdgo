package utils

import "net/http"

func ClickedHyperlink(r *http.Request) string {
	basepath := "/var/semdgo/content"
	link := r.URL.Path
	if link == "/" {
		path := basepath + "/README.md"
		return path
	}
	path := basepath + link
	return path
}
