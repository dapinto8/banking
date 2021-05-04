package main

import (
	"github.com/dapinto8/banking/app"
	"github.com/dapinto8/banking/logger"
)

func main() {
	logger.Info("Starting app...")
	app.Start()
}
