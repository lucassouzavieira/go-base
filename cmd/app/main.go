package main

import (
	web "github.com/lucassvieira/go-project-layout/internal/http"
	logging "github.com/lucassvieira/go-project-layout/internal/logging"
)

func main() {
	logging.LogMessage("Logging out...")
	web.InitHttpServer()
}
