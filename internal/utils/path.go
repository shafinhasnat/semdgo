package utils

import "net/http"

func ClickedHyperlink(r *http.Request, basepath string) string {
	link := r.URL.Path
	if link == "/" {
		return basepath + "/README.md"
	}
	return basepath + link
}
