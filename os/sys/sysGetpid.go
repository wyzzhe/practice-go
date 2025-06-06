package main

import (
	"fmt"
	"syscall"
)

func main() {
	pid := syscall.Getpid()
	fmt.Printf("Current process ID: %d\n", pid)
}
