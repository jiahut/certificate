package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

type twoWayHander struct {
}

func (h *twoWayHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi jazz, this is an example of two way https server")
}

func main() {
	pool := x509.NewCertPool()
	caCrtPath := "ca.crt"
	caCrt, err := ioutil.ReadFile(caCrtPath)
	if err != nil {
		fmt.Println("ReadFile err: ", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":8082",
		Handler: &twoWayHander{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	err = s.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}
