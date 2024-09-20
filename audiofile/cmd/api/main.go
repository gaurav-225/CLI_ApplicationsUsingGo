package main

import (
	"audiofile/services/metadata"
	"flag"
	"fmt"
)


func main() {

	var port int

	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()

	fmt.Println("Starting server on port", port)

	metadata.Run(port)


}


