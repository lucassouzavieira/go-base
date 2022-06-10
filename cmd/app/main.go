package main

import (
	web "github.com/lucassouzavieira/go-base/internal/http"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Logging out something...")
	web.InitHttpServer()
}
