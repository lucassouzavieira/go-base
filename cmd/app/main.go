package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "Logging out..")
	fmt.Println("Writing...")
}
