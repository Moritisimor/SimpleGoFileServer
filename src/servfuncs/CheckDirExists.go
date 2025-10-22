package servfuncs

import "os"

func DirExists(directory string) bool {
	_, err := os.Stat(directory)
	return err == nil
}