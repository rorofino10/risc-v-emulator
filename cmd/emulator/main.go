package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rorofino10/risc-v-emulator/internal/cpu"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	clock := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		processor := cpu.New()

		defer wg.Done()
		processor.Run(clock)
	}()

	ticker := time.NewTicker(1000 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			clock <- struct{}{}
		case <-stop:
			fmt.Println("Stopping")
			close(clock)
			wg.Wait()
			return
		}
	}

}
