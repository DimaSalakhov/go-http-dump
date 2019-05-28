package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

// DefaultPort is the default port to use if once is not specified by the SERVER_PORT environment variable
const DefaultPort = "7893"

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}

	return DefaultPort
}

// Handler dumps incoming requests
func Handler(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(fmt.Sprint("./", time.Now()), requestDump, 0644)

	w.WriteHeader(200)
}

func main() {

	log.Println("starting server, listening on port " + getServerPort())

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":"+getServerPort(), nil)
}
