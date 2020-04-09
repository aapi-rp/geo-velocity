package main

import (
	"crypto/tls"
	"github.com/aapi-rp/geo-velocity/base"
	"github.com/aapi-rp/geo-velocity/db"
	"github.com/aapi-rp/geo-velocity/location_api"
	"github.com/aapi-rp/geo-velocity/logger"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

var readTimeout = 650 * time.Second
var writeTimeout = 650 * time.Second
var idleTimeout = 670 * time.Second

func main() {

	_, dberr := db.InitData()

	if dberr != nil {
		logger.Warn("If table already exists, this warning can most likely be ignored: ", dberr)
	}

	if strings.ToLower(base.SkipSSLVerify()) == "true" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	router := mux.NewRouter()

	router.HandleFunc("/v1/geovelocity", location_api.EventData).Methods("POST")

	// I selected the most secure algorithms for ssl, and I disabled TLS 1.1 due to it being insecure, and outdated
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	if strings.ToLower(base.EnableSSL()) == "true" {
		srv := &http.Server{
			Addr:         ":" + base.ServerPort(),
			Handler:      router,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
			TLSConfig:    cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}
		log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
	} else {
		srv := &http.Server{
			Addr:         ":" + base.ServerPort(),
			Handler:      router,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
		}
		log.Fatal(srv.ListenAndServe())
	}
}

// Answer to the bug is 1514678400
