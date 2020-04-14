package spideCloud

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"time"
)

func PingCdn(cdn *CdnRegion) {
	pinger, err := ping.NewPinger(cdn.domain)
	if err != nil {
		fmt.Println(*cdn, err)
		return
	}

	pinger.Count = Count
	pinger.Interval = time.Duration(Interval) * time.Millisecond
	pinger.Timeout = time.Duration(Count*200+2000) * time.Millisecond
	pinger.SetPrivileged(true)
	pinger.Run()

	stats := pinger.Statistics()
	cdn.rtt = stats.AvgRtt
	cdn.loss = stats.PacketLoss
}
