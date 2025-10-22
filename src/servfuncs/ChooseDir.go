package servfuncs

import (
	"SimpleGoFileServer/colorprint"
	"fmt"
	"os"
	"strings"
)

func ChooseDir(checkee string, directory *string) {
	if strings.TrimSpace(checkee) == "" {
		colorprint.PrintRed("[ ERR ] Directory may not be empty.")
		os.Exit(1)
	}

	if !DirExists(checkee) {
		colorprint.PrintRed(fmt.Sprintf("[ ERR ] '%s' is not a valid directory", checkee))
		os.Exit(1)
	}

	*directory = strings.TrimSpace(checkee)
}
