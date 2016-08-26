package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	pool := x509.NewCertPool()
	caCrtPath := "ca.crt"
	caCrt, err := ioutil.ReadFile(caCrtPath)
	if err != nil {
		fmt.Println("ReadCACrtFile err:", err)
	}
	pool.AppendCertsFromPEM(caCrt)

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://server.zhangzhijia.io:8081")
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}
