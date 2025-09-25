package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Check if a string is worthy of being a port number, and, if so, changes the original string to be that port number
func changePort(checkee string, port *string) {
	checkeeAsInt, err := strconv.ParseInt(checkee, 0, 32)
	// Only parse to int to see if the string is a number at all and to check if it is in the valid range of numbers

	if err != nil {
		fmt.Println("[ ERR ] Ports may only be numbers ranging from 1-65535")
		os.Exit(1)
	}

	if checkeeAsInt < 1 {
		fmt.Println("[ ERR ] Selected port may not be smaller than 1.")
		os.Exit(2)
	} else if checkeeAsInt > 65535 {
		fmt.Println("[ ERR ] Selected port may not be bigger than 65535")
		os.Exit(2)
	}

	*port = checkee
}

// Check if the 2nd argument is a boolean value, and, if so, change the original string accordingly.
func changeHost(localhost string, host *string) {
	if normalize(localhost) == "false" {
		*host = "0.0.0.0:" // Listen on all interfaces
	} else if normalize(localhost) == "true" {
		*host = "localhost:"
	} else {
		fmt.Println("[ ERR ] 2nd Value may only be true or false.")
		os.Exit(3)
	}
}

func changeDir(checkee string, directory *string) {
	if normalize(checkee) == "" {
		fmt.Println("[ ERR ] Directory may not be empty.")
		os.Exit(5)
	}

	checkDirExists(checkee)

	*directory = normalize(checkee)
}

func checkDirExists(directory string) {
	_, err := os.Stat(directory)
	if err != nil {
		fmt.Printf("[ ERR ] The specified directory does not exist: %s\n", directory)
		os.Exit(6)
	}
}

func normalize(unnormalString string) string {
	return strings.ToLower(strings.TrimSpace(unnormalString))
}

func main() {
	port := "8000"
	host := "0.0.0.0:"
	directory := "files"

	switch len(os.Args) {
	case 1:
		checkDirExists("files")
		fmt.Println("[ INFO ] No args, standarding port to 8000, serving files in 'files' directory and " +
			"listening on all interfaces.")

	case 2:
		checkDirExists("files")
		changePort(os.Args[1], &port)

		fmt.Printf("[ INFO ] Port defined, serving files in 'files' directory and listening on "+
			"port %s on all interfaces.\n", port)

	case 3:
		checkDirExists("files")
		changePort(os.Args[1], &port)
		changeHost(os.Args[2], &host)

		fmt.Printf("[ INFO ] Port defined, serving files in 'files' directory and listening on port %s ", port)
		if host == "0.0.0.0:" {
			fmt.Println("on all interfaces.")
		} else {
			fmt.Println("on localhost.")
		}

	case 4:
		changePort(os.Args[1], &port)
		changeHost(os.Args[2], &host)
		changeDir(os.Args[3], &directory)

		fmt.Printf("[ INFO ] Port defined, serving files in '%s' directory and listening "+
			"on port %s ", directory, port)
		if host == "0.0.0.0:" {
			fmt.Println("on all interfaces.")
		} else {
			fmt.Println("on localhost.")
		}
	}

	if len(os.Args) > 4 {
		fmt.Println("[ ERR ] Expected 0-3 Arguments, no more.")
		os.Exit(4)
	}

	fmt.Println("[ INFO ] Starting Server...")
	fmt.Printf("[ INFO ] Listening on Port %s.\n", port)
	http.Handle("/", http.FileServer(http.Dir(directory)))
	fmt.Printf("[ INFO ] Files are available in the %s directory.\n", directory)
	listeningErr := http.ListenAndServe(host+port, nil)

	if listeningErr != nil {
		fmt.Printf("[ ERR ] Something went wrong and the server will shut down now.\nError: %e\n", listeningErr)
	}
}
