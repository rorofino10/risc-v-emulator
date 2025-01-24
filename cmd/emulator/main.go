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

	var wg sync.WaitGroup

	processor := cpu.New()
	wg.Add(1)
	go func() {

		defer wg.Done()
		processor.Run()
	}()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			processor.Clock <- struct{}{}
		case <-stop:
			fmt.Println("Stopping")
			close(processor.Clock)
			wg.Wait()
			return
		}
	}

}
