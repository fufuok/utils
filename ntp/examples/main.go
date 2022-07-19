package main

import (
	"context"
	"log"
	"time"

	"github.com/fufuok/utils/ntp"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	clockOffsetChan := ntp.ClockOffsetChan(ctx, 20*time.Second)
	timeChan := ntp.TimeChan(ctx, 15*time.Second, "ntp.ntsc.ac.cn", "time.nist.gov")
	for {
		select {
		case dur := <-clockOffsetChan:
			log.Printf("clock offset: %s, now: %s\n", dur.String(), time.Now().Add(dur))
		case t := <-timeChan:
			log.Printf("now: %s\n", t)
		case <-ctx.Done():
			log.Println("done.")
			return
		}
	}
}
