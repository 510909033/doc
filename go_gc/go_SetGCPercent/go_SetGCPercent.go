package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"
)

func getDir() string {
	//todo
	return "e:/bbt/go_SetGCPercent.log"
}

func monitor() {
	ticker := time.NewTicker(time.Millisecond * 500)
	f, err := os.OpenFile(getDir(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 755)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case <-ticker.C:
			gcStats := &debug.GCStats{}
			/*
				// len(stats.PauseQuantiles) is 5, it will be filled with the minimum,
				// 25%, 50%, 75%, and maximum pause times.
			*/
			gcStats.PauseQuantiles = make([]time.Duration, 5)
			debug.ReadGCStats(gcStats)
			msg := fmt.Sprintf("%+v\n\n", gcStats)
			f.WriteString(msg)
			//log.Println(msg)
		}
	}
}

func main() {

	go monitor()
	//debug.SetGCPercent(-1)
	go func() {
		for {
			b := testSlice()
			_ = b
		}
	}()

	select {}
}

//go:noinlie
func testSlice() []*int {
	var a []*int
	for i := 0; i < 10000000; i++ {
		a = append(a, &i)
	}
	log.Println("once")
	return a
}
