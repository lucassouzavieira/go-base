package main

import (
	logging "github.com/lucassvieira/go-project-layout/internal/logging"
	web "github.com/lucassvieira/go-project-layout/internal/http"
)

func main() {
	logging.LogMessage("Logging out...")
	web.InitHttpServer()
}
