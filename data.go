package main

import (
	"flag"
	"time"
)

type MAGCOLOREFFECT struct {
	Transform [5][5]float32
}

const VK_ESCAPE = 0x1B

var exitTime = flag.Int("exittime", 0, "The effect will automatically stop after this duration (in seconds)")
var noEscape = flag.Bool("noescape", false, "Prevent stopping the effect via the Escape key")
var speed = flag.String("speed", "normal", "Set speed of the effect (slower/slow/normal/fast/fastest)")

var speeds = map[string]time.Duration{
	"slower":  1800 * time.Millisecond,
	"slow":    1200 * time.Millisecond,
	"normal":  1000 * time.Millisecond,
	"fast":    700 * time.Millisecond,
	"fastest": 500 * time.Millisecond,
}
