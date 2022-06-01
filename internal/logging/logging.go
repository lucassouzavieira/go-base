// Example internal package
package logging

import (
	"log"
	"os"
)

func LogMessage(message string) {
	logger := log.New(os.Stdout, "App info: ", 0)
	logger.Println(message)
}
