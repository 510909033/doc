package internal

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func DemoPprofStart() {
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
}
