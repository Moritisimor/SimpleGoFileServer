package colorprint

import "fmt"

func PrintRed(printee string) {
	fmt.Print("\033[31m" + printee + "\033[0m")
}

func PrintBlue(printee string) {
	fmt.Print("\033[34m" + printee + "\033[0m")
}

func PrintGreen(printee string) {
	fmt.Print("\033[32m" + printee + "\033[0m")
}

func PrintYellow(printee string) {
	fmt.Print("\033[33m" + printee + "\033[0m")
}
