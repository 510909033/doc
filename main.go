package main

import (
	"github.com/510909033/doc/go_nocopy"
	"log"
	"reflect"
	"unsafe"
)

func main() {
	//go_pprof.Open()
	//go_gc.Open()
	//a1()
	go_nocopy.OpenNoCopy()
}

func a1() {
	buf := make([]byte, 100)
	for k := range buf {
		buf[k] = byte(k)
	}
	a := buf[:10] // len(a) == 10, cap(a) == 100.
	b := a[:100]  // is valid, since cap(a) == 100.

	log.Println(a)
	log.Println(b)
	//2022/07/14 11:52:08 [0 1 2 3 4 5 6 7 8 9]
	//2022/07/14 11:52:08 [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
}
func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func s2b(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}
