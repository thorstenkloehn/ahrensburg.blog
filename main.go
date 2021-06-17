package main

import (
	"crypto/tls"
	"github.com/robfig/cron/v3"
	"github.com/thorstenkloehn/ahrensburg.digital/module/Rechsliche_Angabe"
	"github.com/thorstenkloehn/ahrensburg.digital/module/wordpress_pages"
	"net/http"
)

func main() {

	c := cron.New()
	c.AddFunc("@hourly", func() { wordpress_pages.Start() })
	c.Start()
	wordpress_pages.Start()
	Rechsliche_Angabe.Start()

	router := http.NewServeMux()
	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("output/wordpress"))))
	router.Handle("/statik/", http.StripPrefix("/statik/", http.FileServer(http.Dir("statik"))))
	router.Handle("/rechtliche_Angabe/", http.StripPrefix("/rechtliche_Angabe/", http.FileServer(http.Dir("output/Rechtliche_Angabe"))))
	router.Handle("/dokumenten/", http.StripPrefix("/dokumenten/", http.FileServer(http.Dir("docs/static"))))
	cfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		},
	}

	srv := &http.Server{
		Addr:      "ahrensburg.digital:443",
		Handler:   router,
		TLSConfig: cfg,
		// Consider setting ReadTimeout, WriteTimeout, and IdleTimeout
		// to prevent connections from taking resources indefinitely.
	}

	err := srv.ListenAndServeTLS("/etc/letsencrypt/live/ahrensburg.digital/cert.pem", "/etc/letsencrypt/live/ahrensburg.digital/privkey.pem")
	if err != nil {
		http.ListenAndServe(":80", router)
	}

}
