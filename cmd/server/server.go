package main

import (
	"hidden_cyberghost/internal"
	"hidden_cyberghost/pkg/logger"
)

func main() {

	log := logger.NewLogger()
	log.Info("starting Hidden CyberGhost server")

	internal.InitHiddenCyberGhostServer(log)
}
