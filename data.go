package main

import (
	"sync"
	"sync/atomic"
)

type MAGCOLOREFFECT struct {
	Transform [5][5]float32
}

var stop atomic.Bool
var ws sync.WaitGroup

const VK_ESCAPE = 0x1B
