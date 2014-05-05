package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"strconv"
)

func main() {
	var megabytes int
	if len(os.Args) < 1 {
		megabytes = 1
	} else {
		megabytes, _ = strconv.Atoi(os.Args[1])
	}
	mem := make([]byte, 1024 * 1024 * megabytes)
	fmt.Println("Megabytes allocated:", megabytes, "(", len(mem),")")
	sc := make(chan os.Signal, 2)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT)
	<-sc
}
