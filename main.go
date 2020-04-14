package main

import (
	"bestCloud/spideCloud"
	"flag"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

func main() {
	flag.Parse()

	cdns := spideCloud.Spider()
	max := len(cdns)
	wg := sync.WaitGroup{}
	wg.Add(max)

	fmt.Printf(".....................cloud ping: regions(%d), count(%d), interval(%dms), max_timeout(%ds), max_select_loss(%d)...................\n", len(cdns), spideCloud.Count, spideCloud.Interval, spideCloud.Timeout, spideCloud.MaxLoss)

	for index := 0; index < max; index++ {
		go func(cdn *spideCloud.CdnRegion) {
			defer wg.Done()
			spideCloud.PingCdn(cdn)
		}(&cdns[index])
	}

	var ch = make(chan bool)
	go func() {
		wg.Wait()
		ch <- false
	}()

	go func() {
		index := 1
		for {
			time.Sleep(time.Second)
			msg := fmt.Sprintf("ping(%ds)", index)
			output := msg + strings.Repeat(".", 14-len(msg))
			fmt.Println(output)
			index = index + 1
			if index >= int(spideCloud.Timeout) {
				break
			}
		}
	}()

	select {
	case <-time.After(time.Second * time.Duration(spideCloud.Timeout)):
		fmt.Printf("....................................................................................................................................\n")
	case <-ch:
		fmt.Printf("....................................................................................................................................\n")
	}

	var filters spideCloud.CdnRegions
	for _, cdn := range cdns {
		if cdn.Allowed() {
			filters = append(filters, cdn)
		}
	}

	sort.Sort(filters)

	fmt.Printf(".......................................select * from regions(%d) when loss < %d order by rtt asc....................................\n", len(filters), int(spideCloud.MaxLoss))
	fmt.Printf("....................................................................................................................................\n")
	fmt.Println(filters)
}
