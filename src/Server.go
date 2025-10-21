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

	for index, arg := range os.Args {
		if strings.HasPrefix(arg, "--") {
			flag := os.Args[index]
			if len(os.Args) < index+1 {
				fmt.Printf("Expected value after flag: %s\n", flag)
				return
			}

			switch flag {
			default:
				fmt.Printf("Unknown flag: %s\n", flag)
				return

			case "--port":
				changePort(os.Args[index+1], &port)

			case "--host":
				host = os.Args[index+1]

			case "--dir":
				changeDir(os.Args[index+1], &directory)
			}
		}
	}

	fmt.Println("[ INFO ] Starting Server...")
	fmt.Printf("[ INFO ] Listening on %s\n", host+":"+port)
	http.Handle("/", http.FileServer(http.Dir(directory)))
	fmt.Printf("[ INFO ] Files are available in the '%s' directory.\n", directory)
	listeningErr := http.ListenAndServe(host+":"+port, nil)

	if listeningErr != nil {
		fmt.Printf("[ ERR ] Something went wrong and the server will shut down now.\nError: %s\n", listeningErr.Error())
	}
}
