package main

import (
	"fmt"
	"hello-go/internal/server"
)

func main() {
	srv := server.NewServer()
	err := srv.Run()
	if err != nil {
		fmt.Printf("Err: %+v", err)
		return
	}
}
