package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {

	// Define flags
	var filedir, addr string
	flag.StringVar(&filedir, "filedir", ".", "the directory for serving files")
	flag.StringVar(&addr, "addr", ":8090", "The local address to bind")
	flag.Parse()

	// Add the file server ontop of the provided file directory
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(filedir))))

	// Start the server
	fmt.Printf("Starting fileserver using bind address: %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Error occurred: %s", err.Error())
		os.Exit(1)
	}
}
