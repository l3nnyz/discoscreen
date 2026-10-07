package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	flag.Parse()
	var errorChannel = make(chan error, 1)
	var timeout <-chan time.Time

	ws.Add(1)
	go effect(&stop, &ws, errorChannel)

	if *exitTime > 0 {
		timeout = time.After(time.Duration(*exitTime) * time.Second)
	}

	for {
		select {
		case err := <-errorChannel:
			if err != nil {
				stop.Store(true)
				ws.Wait()
				fmt.Println(err)
				os.Exit(1)
			}
		case <-timeout:
			stop.Store(true)
			ws.Wait()
			fmt.Println("Time limit reached. Closing...")
			os.Exit(0)
		default:
			if isEscape() && !*canEscape {
				stop.Store(true)
				ws.Wait()
				fmt.Println("Escape pressed. Closing...")
				os.Exit(0)
			}
		}
		time.Sleep(60 * time.Millisecond)
	}
}

func effect(stop *atomic.Bool, g *sync.WaitGroup, errChannel chan<- error) {
	defer g.Done()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	_, _, err := MagInitialize.Call()
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		errChannel <- fmt.Errorf("MagInitialize call error: %s", err)
		return
	}

	defer MagUninitialize.Call()

	for !stop.Load() {
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
			effect.Transform[4][i] = randVal*0.05 - 0.025
		}

		_, _, err := MagSetFullscreenColorEffect.Call(uintptr(unsafe.Pointer(&effect)))
		if !errors.Is(err, windows.ERROR_SUCCESS) && !errors.Is(err, windows.ERROR_NOT_READY) {
			errChannel <- fmt.Errorf("MagSetFullscreenColorEffect call error: %s", err)
			return
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func isEscape() bool {
	ret, _, _ := GetAsyncKeyState.Call(uintptr(VK_ESCAPE))
	return int16(ret) < 0
}
