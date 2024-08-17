package internal

import (
	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
	"hidden_cyberghost/pkg/logger"
	"os/exec"
)

func InitHiddenCyberGhostServer(log *logger.Logger) {
	cfg := water.Config{
		DeviceType: water.TUN,
	}

	cfg.Name = "hcg"

	ifce, err := water.New(cfg)
	if err != nil {
		log.Error("fail to run hidden cyber ghost server. err => %v", err.Error())
		return
	}

	log.Debug("hidden cyber ghost server started. %v", ifce.Name())

	var frame ethernet.Frame

	addIpCommand := exec.Command("ip", "addr", "add", "10.1.0.10/24", "dev", "hcg")
	addIpResultErr := addIpCommand.Run()
	if addIpResultErr != nil {
		log.Error("fail to run add ipr. %v", addIpResultErr)
	}

	linkIpCommand := exec.Command("ip", "link", "set", "dev", "hcg", "up")
	linkIpResultErr := linkIpCommand.Run()
	if linkIpResultErr != nil {
		log.Error("fail to run link ip. %v", linkIpResultErr)
	}

	for {
		frame.Resize(1500)
		n, err := ifce.Read([]byte(frame))
		if err != nil {
			log.Error(err.Error())
		}
		frame = frame[:n]

		log.Debug("Dst: %s", frame.Destination())
		log.Debug("Src: %s", frame.Source())
		log.Debug("protocol %v", frame.Ethertype())
		log.Debug("Ethertype: % x", frame.Ethertype())
		log.Debug("Payload: % x", frame.Payload())
	}
}
