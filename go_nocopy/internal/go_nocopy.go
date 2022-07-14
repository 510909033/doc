package internal

import (
	"log"
)

type noCopy struct{}    //nolint:unused
func (*noCopy) Lock()   {} //nolint:unused
func (*noCopy) Unlock() {} //nolint:unused

type User struct {
	noCopy
	Id int
}

// User 结构体增加noCopy后， 如果复制了这个变量， go vet能够检测出来
func DemoNoCopy() {
	/*
		go vet ./...
		# github.com/510909033/doc/go_nocopy/internal
		go_nocopy/internal/go_nocopy.go:19:27: call of log.Printf copies lock value: github.com/510909033/doc/go_nocopy/internal.User
		go_nocopy/internal/go_nocopy.go:20:27: call of log.Printf copies lock value: github.com/510909033/doc/go_nocopy/internal.User
	*/
	user := User{}
	user.Id = 1

	log.Printf("user=%+v\n", user)
	log.Printf("user=%+v\n", user)

}
