package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	setLogger()
	http.HandleFunc("/health", ServeHealth)
	if os.Getenv("TLS") == "true" {
		cert := "/etc/admission-webhook/tls/tls.crt"
		key := "/etc/admission-webhook/tls/tls.key"
		logrus.Print("Listening on port 443...")
		logrus.Fatal(http.ListenAndServeTLS(":443", cert, key, nil))
	} else {
		logrus.Print("Listening on port 8080...")
		logrus.Fatal(http.ListenAndServe(":8080", nil))
	}

}

// ServeHealth returns 200 when things are good
func ServeHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("function\n")
	logrus.WithField("uri", r.RequestURI).Debug("healthy")
	fmt.Fprint(w, "OK")
}

func setLogger() {
	logrus.SetLevel(logrus.DebugLevel)

	lev := os.Getenv("LOG_LEVEL")
	if lev != "" {
		llev, err := logrus.ParseLevel(lev)
		if err != nil {
			logrus.Fatalf("cannot set LOG_LEVEL to %q", lev)
		}
		logrus.SetLevel(llev)
	}

	if os.Getenv("LOG_JSON") == "true" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}
