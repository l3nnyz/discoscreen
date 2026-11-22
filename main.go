package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	flag.Parse()
	var errorChannel = make(chan error, 1)
	var timeout <-chan time.Time
	var delay time.Duration
	setSpeed(&delay)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ws.Add(1)
	go effect(stop, &ws, &delay, errorChannel)

	if *exitTime > 0 {
		timeout = time.After(time.Duration(*exitTime) * time.Second)
	}

	printHelp()

	start := time.Now()
	lastShown := time.Duration(-1)

	for {
		select {
		case err := <-errorChannel:
			if err != nil {
				shutdown(err.Error(), 1)
			}
		case <-timeout:
			shutdown("Time limit reached. Closing...", 0)
		case <-interrupt:
			shutdown("Ctrl+C pressed. Closing...", 0)
		default:
			if isEscape() && !*noEscape {
				shutdown("Escape pressed. Closing...", 0)
			}
		}

		if elapsed := time.Since(start).Truncate(time.Second); elapsed != lastShown {
			lastShown = elapsed
			fmt.Printf("\rRunning time: %s", formatDuration(elapsed))
		}

		time.Sleep(60 * time.Millisecond)
	}
}

func printHelp() {
	if *noEscape {
		fmt.Println("Press Ctrl+C to stop the effect and close the program.")
	} else {
		fmt.Println("Press Escape or Ctrl+C to stop the effect and close the program.")
	}
	if *exitTime > 0 {
		fmt.Printf("The effect will stop automatically after %s.\n", formatDuration(time.Duration(*exitTime)*time.Second))
	}
}

func shutdown(message string, code int) {
	close(stop)
	ws.Wait()
	fmt.Println()
	fmt.Println(message)
	os.Exit(code)
}

func formatDuration(d time.Duration) string {
	total := int(d / time.Second)
	return fmt.Sprintf("%02d:%02d:%02d", total/3600, total%3600/60, total%60)
}

func effect(stop <-chan struct{}, g *sync.WaitGroup, delay *time.Duration, errChannel chan<- error) {
	defer g.Done()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	_, _, err := MagInitialize.Call()
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		errChannel <- fmt.Errorf("MagInitialize call error: %s", err)
		return
	}

	defer MagUninitialize.Call()

	for {
		effect := MAGCOLOREFFECT{
			Transform: [5][5]float32{
				{1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 0},
				{0, 0, 0, 0, 1},
			},
		}

		for i := 0; i < 3; i++ {
			randVal := float32(rand.IntN(0xFFFFFF)) / float32(0xFFFFFF)
			effect.Transform[4][i] = randVal*0.05 - 0.035
		}

		_, _, err := MagSetFullscreenColorEffect.Call(uintptr(unsafe.Pointer(&effect)))
		if !errors.Is(err, windows.ERROR_SUCCESS) && !errors.Is(err, windows.ERROR_NOT_READY) {
			errChannel <- fmt.Errorf("MagSetFullscreenColorEffect call error: %s", err)
			return
		}

		select {
		case <-stop:
			return
		case <-time.After(*delay):
		}
	}
}

func isEscape() bool {
	ret, _, _ := GetAsyncKeyState.Call(uintptr(VK_ESCAPE))
	return int16(ret) < 0
}

func setSpeed(delay *time.Duration) {
	switch *speed {
	case "slower":
		*delay = slowerSpeed
	case "slow":
		*delay = slowSpeed
	case "normal":
		*delay = normalSpeed
	case "fast":
		*delay = fastSpeed
	case "fastest":
		*delay = fastestSpeed
	case "club":
		*delay = clubSpeed
	default:
		*delay = normalSpeed
	}
}
