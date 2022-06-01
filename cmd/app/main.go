package main

import (
	web "github.com/lucassouzavieira/go-base/internal/http"
	logging "github.com/lucassouzavieira/go-base/internal/logging"
)

func main() {
	logging.LogMessage("Logging out...")
	web.InitHttpServer()
}
