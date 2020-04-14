package spideCloud

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

var Count int
var Interval int64
var Timeout int64
var MaxLoss int64

func init() {
	flag.IntVar(&Count, "c", 10, "ping count")
	flag.Int64Var(&Interval, "i", 200, "ping interval(ms)")
	flag.Int64Var(&Timeout, "t", 2, "ping timeout(s)")
	flag.Int64Var(&MaxLoss, "m", 5, "max ping loss")
}

type CdnRegion struct {
	name   string
	region string
	domain string
	rtt    time.Duration
	loss   float64
}

func (cdn CdnRegion) Allowed() bool {
	if cdn.loss >= 0 && cdn.loss < float64(MaxLoss) && cdn.rtt > 0 {
		return true
	}
	return false
}

func (cdn CdnRegion) String() string {
	translate := func(s string, n int) string {
		size := (len(s) - len([]rune(s))) / 2
		blanks := n - len(s) + size
		if blanks > 0 {
			return s + strings.Repeat(" ", blanks)
		} else {
			//fmt.Println(s)
		}
		return s
	}
	name := translate(cdn.name, 20)
	region := translate(cdn.region, 20)
	domain := translate(cdn.domain, 80)
	return fmt.Sprintf("%s%s%s%4dms\t%4d\n", name, region, domain, cdn.rtt/time.Millisecond, int(cdn.loss))
}

type CdnRegions []CdnRegion

func (s CdnRegions) Len() int           { return len(s) }
func (s CdnRegions) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s CdnRegions) Less(i, j int) bool { return s[i].rtt < s[j].rtt }
