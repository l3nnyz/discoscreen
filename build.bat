@echo off
go build -ldflags="-s -w -H=windowsgui" -trimpath -o discoscreen.exe
