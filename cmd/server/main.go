package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"

	"github.com/shafinhasnat/semdgo/internal/config"
	"github.com/shafinhasnat/semdgo/internal/handler"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	http.HandleFunc("/", handler.Handle)

	if cfg.LetsEncrypt.Enabled {
		if cfg.LetsEncrypt.Domain == "" {
			log.Fatal("letsencrypt.domain must be set when letsencrypt.enabled is true")
		}

		m := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(cfg.LetsEncrypt.Domain),
			Cache:      autocert.DirCache("/var/semdgo/certs"),
		}
		if cfg.LetsEncrypt.Email != "" {
			m.Email = cfg.LetsEncrypt.Email
		}

		// HTTP -> HTTPS redirect
		go func() {
			log.Println("Starting HTTP redirect server on :80")
			if err := http.ListenAndServe(":80", m.HTTPHandler(nil)); err != nil {
				log.Printf("HTTP redirect server error: %v", err)
			}
		}()

		log.Printf("Starting HTTPS server on :443 for domain %s", cfg.LetsEncrypt.Domain)
		server := &http.Server{
			Addr:      ":443",
			TLSConfig: m.TLSConfig(),
		}
		log.Fatal(server.ListenAndServeTLS("", ""))
	} else {
		addr := fmt.Sprintf(":%d", cfg.Port)
		log.Printf("Starting server on port %d", cfg.Port)
		log.Fatal(http.ListenAndServe(addr, nil))
	}
}
