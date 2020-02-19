package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	certFile = flag.String("cert", "C:/Users/tgogoi/exercise/go-workspace/src/learning/tls/client/files/localhost-client.cert", "A PEM eoncoded certificate file.")
	keyFile  = flag.String("key", "C:/Users/tgogoi/exercise/go-workspace/src/learning/tls/client/files/localhost-client.key", "A PEM encoded private key file.")
	caFile   = flag.String("CA", "C:/Users/tgogoi/exercise/go-workspace/src/learning/tls/server/files/localhost-server.cert", "A PEM eoncoded CA's certificate file.")
)

func main() {
	flag.Parse()

	// Load client cert
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(*caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup TLS config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		//ServerName: "localhost-server",
	}
	//tlsConfig.BuildNameToCertificate()

	//Simple tls dial client
	_, err = tls.Dial("tcp", "localhost:8443", tlsConfig)
	if err != nil {
		log.Fatal(err)
	}


	// HTTPS client
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Do GET something
	resp, err := client.Get("https://localhost:8443/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}