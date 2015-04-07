//Web and socker server to generate UUID5 for Global Names
package main

import (
	"fmt"
	"os"

	"github.com/gnames/uuid5"
)

func main() {
	args := os.Args
	if len(args) != 2 || !(args[1] == "web" || args[1] == "socket") {
		usage(args[0])
		return
	}
	uuid := uuid5.UUID5("hello")
	fmt.Println(args)
	fmt.Println(uuid.String())
}

func usage(cmd string) {
	fmt.Println("Usage:")
	fmt.Printf("\n%s web [-port 12345] [-domain example.com]\n", cmd)
	fmt.Println("  Defaults: -port 8080 -domain localhost")
	fmt.Print("\nor\n\n")
	fmt.Printf("%s socket [-port 12345]\n", cmd)
	fmt.Println("  Defaults -port 4335")
}
