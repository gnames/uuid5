Global Names UUID5 Generator
============================

It is a simple web server which takes strings and creates a corresponding
UUID version 5 using DNS globalnames.org as a namespace. More information about
UUID version 5 creation can be found at [RFC-4122][1]

Compiling
---------

1. Setup standard Go environment http://golang.org/doc/code.html and ensure that $GOPATH environment variable properly set.
2. `go get github.com/GlobalNamesArchitecture/uuid5`.
3. `cd $GOPATH/src/github.com/GlobalNamesArchitecture/uuid5`
4. `go build` to get binary

Usage
-----
    ./uuid5 --port 80 --url http://your_url



[1]: http://www.ietf.org/rfc/rfc4122.txt
