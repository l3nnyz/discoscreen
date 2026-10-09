package main

import (
	"flag"
	"sync"
	"sync/atomic"
	"time"
)

type MAGCOLOREFFECT struct {
	Transform [5][5]float32
}

var stop atomic.Bool
var ws sync.WaitGroup

const VK_ESCAPE = 0x1B

var exitTime = flag.Int("exittime", 0, "The effect will automatically stop after this duration (in seconds)")
var canEscape = flag.Bool("noescape", false, "Prevent stopping the effect via the Escape key")
var speed = flag.String("speed", "normal", "Set speed of the effect (slow/normal/fast/fastest)")

const (
	slowerSpeed  time.Duration = 1800
	slowSpeed    time.Duration = 1200
	normalSpeed  time.Duration = 1000
	fastSpeed    time.Duration = 700
	fastestSpeed time.Duration = 500
	clubSpeed    time.Duration = 100
)
