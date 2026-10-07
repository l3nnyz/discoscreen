package main

import (
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func main() {
	ws.Add(1)
	go effect(&stop, &ws)

	for !isEscape() {
		time.Sleep(50 * time.Millisecond)
	}
	stop.Store(true)

	ws.Wait()
	os.Exit(0)
}

func effect(stop *atomic.Bool, g *sync.WaitGroup) {
	MagInitialize.Call()
	defer MagUninitialize.Call()
	defer g.Done()

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
			randVal := float32(rand.Intn(0xFFFFFF)) / float32(0xFFFFFF)
			effect.Transform[4][i] = randVal*0.05 - 0.025
		}

		MagSetFullscreenColorEffect.Call(uintptr(unsafe.Pointer(&effect)))

		time.Sleep(500 * time.Millisecond)
	}
}

func isEscape() bool {
	ret, _, _ := GetAsyncKeyState.Call(uintptr(VK_ESCAPE))
	return int16(ret) < 0
}
