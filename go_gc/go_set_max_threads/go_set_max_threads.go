package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"runtime/debug"

	"time"
)

func main() {

	maxThreads := debug.SetMaxThreads(100)
	log.Println("before maxThreads=", maxThreads)

	for i := 0; i < 110; i++ {
		go func(i int) {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()

			time.Sleep(time.Minute)
			log.Println("i=", i)
		}(i)
	}

	go func() {
		//var p []runtime.StackRecord
		var p = make([]runtime.StackRecord, 100)
		for {
			n, ok := runtime.ThreadCreateProfile(p)
			if !ok {
				log.Println("! ok， n=", n)
				return
			}
			log.Println("len(p)=", len(p), "n=", n)
			for k := range p {
				//大概就是这么用的
				sp := p[k]
				runtime.CallersFrames(sp.Stack())
			}
			time.Sleep(time.Second * 5)
		}
	}()
	//runtime.NumGoroutine()
	//http://localhost:8008/debug/pprof/threadcreate?debug=1
	log.Fatal(http.ListenAndServe(":8008", nil))
	select {}
	//runtime.KeepAlive()
}
