package common

import (
	"fmt"
	"os"
)

// QuitIfError prints the message and error and exits if the error is not nil
func QuitIfError(err error, message string) {
	if err == nil {
		return
	}
	fmt.Println(message, err)
	os.Exit(1)
}
