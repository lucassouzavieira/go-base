package main

import (
	web "github.com/lucassvieira/go-base/internal/http"
	logging "github.com/lucassvieira/go-base/internal/logging"
)

func main() {
	logging.LogMessage("Logging out...")
	web.InitHttpServer()
}
