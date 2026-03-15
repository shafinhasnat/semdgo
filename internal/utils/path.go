package utils

import (
	"net/http"
	"path/filepath"
)

func ClickedHyperlink(r *http.Request, basepath, entrypoint string) string {
	link := r.URL.Path
	if link == "/" {
		return filepath.Join(basepath, entrypoint)
	}
	return filepath.Join(basepath, link)
}
