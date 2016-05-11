package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {

	// Define flags
	var filedir, addr string
	flag.StringVar(&filedir, "filedir", ".", "the directory for serving files")
	flag.StringVar(&addr, "addr", ":8090", "The local address to bind")
	flag.Parse()

	// Add the file server for the asset directory
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(filedir))))

	// Start the server
	fmt.Printf("Starting server using address: %s", addr)
	http.ListenAndServe(addr, nil)
}
