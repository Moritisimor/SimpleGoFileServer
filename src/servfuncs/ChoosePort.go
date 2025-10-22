package servfuncs

import (
	"SimpleGoFileServer/colorprint"
	"os"
	"strconv"
)

func ChoosePort(checkee string, port *string) {
	checkeeAsInt, err := strconv.ParseInt(checkee, 0, 32)
	// Only parse to int to see if the string is a number at all and to check if it is in the valid range of numbers

	if err != nil {
		colorprint.PrintRed("[ ERR ] Ports may only be numbers ranging from 1-65535")
		os.Exit(1)
	}

	if checkeeAsInt < 1 {
		colorprint.PrintRed("[ ERR ] Selected port may not be smaller than 1.")
		os.Exit(1)
	} 

	if checkeeAsInt > 65535 {
		colorprint.PrintRed("[ ERR ] Selected port may not be bigger than 65535")
		os.Exit(1)
	}

	if checkeeAsInt < 1024 {
		colorprint.PrintYellow("[ WARN ] Be careful when using Ports below 1024 as those are reserved.\n")
	}

	*port = checkee
}