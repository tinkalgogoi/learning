package main

import (
	"net/http"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":8443", "C:/Users/tgogoi/exercise/go-workspace/src/learning/tls/server/files/localhost-server.cert", "C:/Users/tgogoi/exercise/go-workspace/src/learning/tls/server/files/localhost-server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}