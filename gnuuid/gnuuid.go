//Web and socker server to generate UUID5 for Global Names
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	server string
	port   int
	domain string
)

func main() {
	processFlags()
	switch server {
	case "socket":
		startSocket()
	case "web":
		startWeb()
	default:
		usage()
	}
}

func processFlags() {
	flag.StringVar(&server, "server", "socket",
		"Server type (web or socket)")
	flag.IntVar(&port, "port", 5445, "Port to access server")
	flag.StringVar(&domain, "domain", "localhost", "Domain of a webserver")
	flag.Parse()
	if len(flag.Args()) > 0 {
		server = ""
	}
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  Start Socket Server on port 5335:")
	fmt.Printf("\n    %s\n\n", os.Args[0])
	fmt.Println("  Start Socket Server on other port:")
	fmt.Printf("\n    %s -port 8888\n\n", os.Args[0])
	fmt.Println("  Start HTTP Server on port 5335, localhost domain:")
	fmt.Printf("\n    %s -server=web\n\n", os.Args[0])
	fmt.Println("  Start Socket Server on other port, domain:")
	fmt.Printf("\n    %s -server=web -port=80 -domain=example.org\n", os.Args[0])
}

func startSocket() {
	fmt.Printf("Starting Socket Server on port %d\n\n", port)
}

func startWeb() {
	fmt.Printf("Starting Web Server on port %d using domain '%s'\n\n",
		port, domain)
}
