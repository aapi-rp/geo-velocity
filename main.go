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
	//geoP, err := utils.GetGeoDataFromIP("91.207.175.104")
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(geoP.LAT, geoP.LONG, geoP.CITY, geoP.SUBDEVISIONNAME)
	//
	//curr := models.Coord{}
	//
	//curr.Longitude = geoP.LONG
	//curr.Latitude = geoP.LAT
	//
	//tuuid := uuid.Must(uuid.NewRandom())
	//
	//geoP.UUID = []byte(tuuid.String())
	//
	//logger.Warn(string(geoP.UUID))
	//
	////models.AddEvents(geoP)
	//
	//geoC, err := utils.GetGeoDataFromIP("206.81.252.6")
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(geoC.LAT, geoC.LONG, geoC.CITY, geoC.SUBDEVISIONNAME)
	//
	//pre := models.Coord{}
	//
	//pre.Longitude = geoC.LONG
	//pre.Latitude = geoC.LAT
	//
	//vdist := utils.VariableDistance(geoP.LAT, geoP.LONG, geoC.LAT, geoC.LONG, "miles")
	//log.Println("Variable Distance: ", vdist)
	//mph := utils.MPH(vdist, 1514764800, 1514851200)
	//
	//endhours, timestamp := utils.GetEndTimeFromDistanceAndSpeed(vdist, float64(mph), 1514851200)
	//starthours, timestamp2 := utils.GetStartTimeFromDistanceAndSpeed(vdist, float64(mph), 1514764800)
	//
	//log.Println("Hours the trip would take: ", endhours, "hours, and ending timestamp of: ", timestamp.Unix())
	//log.Println("Hours the trip would take: ", starthours, "hours and starting timestamp of: ", timestamp2.Unix())
	//log.Println("Miles Per Hour: ", mph)
	//
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
