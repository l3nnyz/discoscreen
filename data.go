package main

import (
	"flag"
	"sync"
	"time"
)

type MAGCOLOREFFECT struct {
	Transform [5][5]float32
}

var stop = make(chan struct{})
var ws sync.WaitGroup

const VK_ESCAPE = 0x1B

var exitTime = flag.Int("exittime", 0, "The effect will automatically stop after this duration (in seconds)")
var noEscape =flag.Bool("noescape", false, "Prevent stopping the effect via the Escape key")
var speed = flag.String("speed", "normal", "Set speed of the effect (slower/slow/normal/fast/fastest)")

const (
	slowerSpeed  = 1800 * time.Millisecond
	slowSpeed    = 1200 * time.Millisecond
	normalSpeed  = 1000 * time.Millisecond
	fastSpeed    = 700 * time.Millisecond
	fastestSpeed = 500 * time.Millisecond
	clubSpeed    = 100 * time.Millisecond
)
