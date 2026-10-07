package main

import (
	"flag"
	"sync"
	"sync/atomic"
)

type MAGCOLOREFFECT struct {
	Transform [5][5]float32
}

var stop atomic.Bool
var ws sync.WaitGroup

const VK_ESCAPE = 0x1B

var exitTime = flag.Int("exittime", 0, "The effect will automatically stop after this duration (in seconds)")
var canEscape = flag.Bool("noescape", false, "Prevent stopping the effect via the Escape key")
