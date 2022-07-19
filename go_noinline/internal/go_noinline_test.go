package internal

import "testing"

//内联 非 内联性能区别
/*
GOROOT=C:\Users\Administrator\go\go1.16.15 #gosetup
GOPATH=C:\Users\Administrator\go #gosetup
C:\Users\Administrator\go\go1.16.15\bin\go.exe test -c -o C:\Users\Administrator\AppData\Local\Temp\___gobench_go_noinline_test_go.exe github.com/510909033/doc/go_noinline/internal #gosetup
C:\Users\Administrator\AppData\Local\Temp\___gobench_go_noinline_test_go.exe -test.v -test.bench "^BenchmarkAddNoinline|BenchmarkAddInline$" -test.run ^$ #gosetup
goos: windows
goarch: amd64
pkg: github.com/510909033/doc/go_noinline/internal
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkAddNoinline
BenchmarkAddNoinline-4   	1000000000	         0.7967 ns/op
BenchmarkAddInline
BenchmarkAddInline-4     	507777457	         2.413 ns/op
PASS
*/

//go:noinline
func AddNoinline(x, y, z int) int {
	return x + y + z
}

func AddInline(x, y, z int) int {
	return x + y + z
}
func BenchmarkAddNoinline(b *testing.B) {
	x, y, z := 1, 2, 3
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddInline(x, y, z)
	}
}

func BenchmarkAddInline(b *testing.B) {
	x, y, z := 1, 2, 3
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddNoinline(x, y, z)
	}
}
