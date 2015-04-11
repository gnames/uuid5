//Web and socker server to generate UUID5 for Global Names
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	server string
	port   string
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
	flag.StringVar(&port, "port", "5445", "Port to access server")
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
	fmt.Printf("Starting Socket Server on port %s\n\n", port)
	sock, err := net.Listen("tcp", "localhost"+":"+port)
	if err != nil {
		fmt.Println("Socket server error: ", err.Error())
		os.Exit(1)
	}
	defer sock.Close()
	socketListen(sock)
}

func startWeb() {
	fmt.Printf("Starting Web Server on port %s using domain '%s'\n\n",
		port, domain)
	http.HandleFunc("/", webHandler)
	http.ListenAndServe(port, nil)
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Path[1:]
	if input == "" {
		webInstructions(w)
	} else {
		webRequest(w, input)
	}
}

func webInstructions(w http.ResponseWriter) {
	fmt.Fprintf(w, "Enter name string in url like \"http://%s/Homo%%20sapiens\"\n", domain)
	fmt.Fprintf(w, "Or enter serversl name strings divided by pipe character like \"http://%s/Homo%%20sapiens%%7CPardosa%%20moesta%%7CParus%%20major%%20(Linnaeus,%%201758)\"\n", domain)
}

func socketListen(sock net.Listener) {
	for {
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println("Input error: ", err.Error())
			os.Exit(1)
		}
		go handleSocket(conn)
	}
}

func handleSocket(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading input: ", err.Error())
	}
	fmt.Println(buf)
	conn.Write([]byte(buf))
}
