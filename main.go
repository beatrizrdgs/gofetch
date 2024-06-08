package main

import (
	"github.com/beatrizrdgs/gofetch/gofetch"
)

const (
	UNKNOWN = "unknown"
)

func main() {
	s := gofetch.NewSystem()
	s.Print()
}
