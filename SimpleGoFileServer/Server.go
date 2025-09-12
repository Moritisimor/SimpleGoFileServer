package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// This function checks if a given string is worthy of being used as a port number, and if not, exits with corresponding exit codes.
func isPortWorthy(checkee string) {
	checkeeAsInt, err := strconv.ParseInt(checkee, 0, 32)

	if err != nil {
		fmt.Println("Ports may only be numbers ranging from 1-65535")
		os.Exit(1)
	}

	if checkeeAsInt < 1 {
		fmt.Println("Selected port may not be smaller than 0.")
		os.Exit(2)
	} else if checkeeAsInt > 65535 {
		fmt.Println("Selected port may not be bigger than 65535")
		os.Exit(2)
	}
}

func normalize(unnormalString string) string {
	return strings.ToLower(strings.TrimSpace(unnormalString))
}

func main() {
	var port 			string
	var localhost 		bool
	var listeningErr	error

	switch len(os.Args) {
	case 1:
		fmt.Println("No args, standarding port to 8000 and listening on all interfaces.")
		port = "8000"
		localhost = false

	case 2:
		isPortWorthy(os.Args[1])
		port = os.Args[1]
		localhost = false
		fmt.Printf("Port defined, listening on %s on all interfaces\n", port)

	case 3:
		isPortWorthy(os.Args[1])
		port = os.Args[1]

		if normalize(os.Args[2]) == "true" {
			localhost = true
			fmt.Printf("Port and localhost defined, listening on %s on localhost.\n", port)
		} else if normalize(os.Args[2]) == "false" {
			localhost = false
			fmt.Printf("Port and localhost defined, listening on %s on all interfaces.\n", port)
		} else {
			fmt.Println("Expected true or false value for 2nd argument.")
			os.Exit(3)
		}

	case 4:
		fmt.Println("Expected 0, 1 or 2 arguments, no more.")
		os.Exit(4)
	}

	fmt.Println("Starting Server...")
	fmt.Printf("Listening on Port %s.\n", port)
	http.Handle("/", http.FileServer(http.Dir("files")))
	fmt.Println("Files are available in the files directory.")

	if localhost {
		listeningErr = http.ListenAndServe("localhost:" + port, nil)
	} else {
		listeningErr = http.ListenAndServe("0.0.0.0:" + port, nil)
	}

	if listeningErr != nil {
		fmt.Printf("Something went wrong and the server will shut down now.\nError: %e\n", listeningErr)
	}
}
