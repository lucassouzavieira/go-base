// Example internal package
package logging

import (
	"log"
	"os"
)

func LogMessage(message string) {
	logger := log.New(os.Stdout, "onboarding-lvieira-info: ", 0)
	logger.Output(2, message)

}
