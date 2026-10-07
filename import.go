package main

import "golang.org/x/sys/windows"

var (
	modMagnification = windows.NewLazySystemDLL("Magnification.dll")

	MagInitialize               = modMagnification.NewProc("MagInitialize")
	MagUninitialize             = modMagnification.NewProc("MagUninitialize")
	MagSetFullscreenColorEffect = modMagnification.NewProc("MagSetFullscreenColorEffect")

	user32 = windows.NewLazySystemDLL("user32.dll")

	GetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)
