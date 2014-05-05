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
	if len(os.Args) < 2 {
		megabytes = 1
	} else {
		megabytes, _ = strconv.Atoi(os.Args[1])
	}
	mem := make([]byte, 1024 * 1024 * megabytes)
	fmt.Println("Allocated:", megabytes, "mb (", len(mem),"b )")
	if len(os.Args) == 3 {
		megabytes, _ := strconv.Atoi(os.Args[2])
		for i := 0; i < (1024 * 1024 * megabytes); i++ {
			mem[i] = byte(255)
		}
		fmt.Println("Used:", megabytes, "mb")
	}
	sc := make(chan os.Signal, 2)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT)
	<-sc
}
