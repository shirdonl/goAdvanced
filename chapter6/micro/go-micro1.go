package main

import (
	"github.com/asim/go-micro"
)

func main() {
	srv := micro.NewService(micro.Name("hello"))

	srv.Run()
}
