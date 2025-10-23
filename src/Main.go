package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"SimpleGoFileServer/servfuncs"
	"SimpleGoFileServer/colorprint"
)

func main() {
	port := "8000"
	host := "0.0.0.0"
	directory := "files"

	for index, arg := range os.Args {
		if strings.HasPrefix(arg, "-") {
			flag := os.Args[index]
			if len(os.Args) < index + 2 {
				colorprint.PrintRed(fmt.Sprintf("Expected value after flag: %s\n", flag))
				return
			}

			switch flag {
			default:
				colorprint.PrintRed(fmt.Sprintf("Unknown flag: %s\n", flag))
				return

			case "--port", "-p":
				servfuncs.ChoosePort(os.Args[index + 1], &port)

			case "--host", "-h":
				host = os.Args[index + 1]

			case "--dir", "-d":
				servfuncs.ChooseDir(os.Args[index + 1], &directory)
			}
		}
	}

	colorprint.PrintBlue("[ INFO ] Starting Server...\n")
	colorprint.PrintBlue(fmt.Sprintf("[ INFO ] Listening on http://%s:%s\n", host, port))
	http.Handle("/", http.FileServer(http.Dir(directory)))

	if !servfuncs.DirExists(directory) {
		colorprint.PrintRed("[ ERR ] 'files' is not a valid directory!\n")
		return
	}
	
	colorprint.PrintBlue(fmt.Sprintf("[ INFO ] Files are available in the '%s' directory.\n", directory))
	listeningErr := http.ListenAndServe(host + ":" + port, nil)

	if listeningErr != nil {
		colorprint.PrintRed(fmt.Sprintf("[ ERR ] Something went wrong and the server will shut down now.\n[ ERR ] Error: %s\n", listeningErr.Error()))
	}
}
